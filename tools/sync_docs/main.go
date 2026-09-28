// sync_docs 把接口文档的字段表与本库的结构体对起来。
//
// 用法（在仓库根目录）：
//
//	go run tools/sync_docs.go -check -docs <路径>   # 只报漂移，有漂移则退出码非 0
//	go run tools/sync_docs.go -write -docs <路径>   # 把文档里有、结构体里没有的字段补上
//	go run tools/sync_docs.go -coverage            # 报告映射覆盖率（还有多少结构体没挂锚点）
//
// `-check` 与 `-write` 必须给 `-docs`，指到接口文档的检出位置；`-coverage` 不读文档，
// 不用给。这个参数没有默认值 —— 那个目录名跟本地布局绑定、跟代码无关，写进仓库会
// 一路带进公开历史。在仓库根目录跑 `make sync-docs` 会自动探测，省掉它。
//
// ## 为什么是「映射表 + 工具」而不是全自动
//
// 文档里的对象叫 `data`、`list[]`、`data`中的`control`对象，而 Go 结构体叫
// CommentsControl、VipUserVip —— **名字对不上**，没法推。靠字段重合度去猜
// 会误配（实测把 `vip_type` 匹到 VipPrivilege、把根对象的 `code` 匹到
// SearchRespData），而**往错的结构体加字段编译器不报错**，会静默上线。
//
// 所以锚点写死在 tools/docmap.go 里：一个结构体挂到文档的哪一节、哪张表。
// 那是人的判断，机器做不了，但做完一次就能一直用。
//
// ## 这个工具能做什么、不能做什么
//
// 能：报「文档有、Go 没有」（新增字段漏了）、报「Go 有、文档没有」（拼错或改名）、
// 按文档补齐字段。
//
// 不能：判断 `num` 该是 int 还是指针还是 bool、判断两个结构体该不该复用、
// 决定新结构体叫什么名字。那些仍然要人看。
package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// docTable 是文档里的一张字段表。
type docTable struct {
	Labels []string // 表前面那几行标签，如 "`data`中的`control`对象"
	Order  []string // 字段出现顺序
	Rows   map[string]docRow
}

type docRow struct {
	Kind    string // num / str / obj / array / bool ...
	Content string
	Note    string
}

// goStruct 是一个 Go 结构体。
type goStruct struct {
	Name   string
	File   string
	Line   int
	Fields []string // json 标签，按出现顺序
	Tags   map[string]bool
}

var (
	flagCheck    = flag.Bool("check", false, "只检查，有漂移则退出码非 0")
	flagWrite    = flag.Bool("write", false, "把缺失字段写回 Go 文件")
	flagCoverage = flag.Bool("coverage", false, "报告映射覆盖率")
	flagDocs     = flag.String("docs", "", "接口文档检出路径（-check / -write 必填）")
	flagQuiet    = flag.Bool("quiet", false, "只打印汇总")
	flagVerbose  = flag.Bool("v", false, "打印每个锚点实际匹到的表标签")
)

func main() {
	flag.Parse()

	if !*flagCheck && !*flagWrite && !*flagCoverage {
		*flagCheck = true
	}

	root, err := os.Getwd()
	must(err)

	structs, err := loadStructs(root)
	must(err)

	if *flagCoverage {
		reportCoverage(root, structs)
		return
	}

	if *flagDocs == "" {
		fmt.Fprintln(os.Stderr, "✗ 没给接口文档检出位置。")
		fmt.Fprintln(os.Stderr, "    用 -docs 指定：go run ./tools/sync_docs -check -docs ../某个目录")
		fmt.Fprintln(os.Stderr, "    或者在仓库根目录跑 make sync-docs（会自动探测检出位置）")
		os.Exit(2)
	}

	drift := 0
	for _, name := range sortedKeys(docMap) {
		locs := docMap[name]

		gs, ok := structs[name]
		if !ok {
			fmt.Printf("✗ %s：docmap 里有，但包里找不到这个结构体\n", name)
			drift++
			continue
		}

		// 一个结构体可能挂多处（`Label` 被多处共用），按并集比对
		table, deadAnchors, err := readTables(*flagDocs, locs)
		if err != nil {
			fmt.Printf("✗ %s：%v\n", name, err)
			drift++
			continue
		}
		if len(deadAnchors) > 0 {
			fmt.Printf("✗ %s：%d 处锚点读不到表（并集因此不完整）\n", name, len(deadAnchors))
			for _, d := range deadAnchors {
				fmt.Printf("      ! %s\n", d)
			}
			drift++
		}

		missing, extra, known := diff(gs, table, knownSet(docKnown[name]))

		// -v 时把匹到的表标签打出来。锚点写错、或者标签撞车匹到别的表上，
		// 从字段的增减看不出来（并集只会变宽），但标签会立刻露馅
		if *flagVerbose {
			fmt.Printf("\n%s  (%s:%d)  %d 个字段\n", name, gs.File, gs.Line, len(gs.Fields))
			for _, l := range locs {
				fmt.Printf("  锚点：%s  ## %s  →  %s\n", l.File, l.Section, l.Label)
			}
			for _, l := range table.Labels {
				fmt.Printf("  匹到：%s\n", firstLine(l))
			}
			if len(known) > 0 {
				fmt.Printf("  docKnown 抵扣 %d 个\n", len(known))
			}
		}

		if len(missing) == 0 && len(extra) == 0 {
			continue
		}
		drift++

		fmt.Printf("\n%s  (%s:%d)\n", name, gs.File, gs.Line)
		for _, l := range locs {
			fmt.Printf("  文档：%s  ## %s  →  %s\n", l.File, l.Section, l.Label)
		}
		if len(missing) > 0 {
			fmt.Printf("  文档有、Go 没有（%d）：\n", len(missing))
			for _, k := range missing {
				r := table.Rows[k]
				fmt.Printf("      + %-30s %-6s %s\n", k, r.Kind, firstLine(r.Content+r.note()))
			}
			if *flagWrite {
				if err := writeFields(root, gs, missing, table); err != nil {
					fmt.Printf("      ✗ 写入失败：%v\n", err)
				} else {
					fmt.Printf("      ✓ 已写入\n")
				}
			}
		}
		if len(extra) > 0 {
			fmt.Printf("  Go 有、文档没有（%d）—— 拼错、已改名，或这个结构体还服务于别的接口？\n", len(extra))
			for _, k := range extra {
				fmt.Printf("      - %s\n", k)
			}
		}
		if len(known) > 0 {
			fmt.Printf("  （另有 %d 个字段记在 tools/docmap.go 的 docKnown 里，已核对）\n", len(known))
		}
	}

	if !*flagQuiet {
		fmt.Println()
	}
	if drift == 0 {
		fmt.Printf("✓ %d 个结构体与文档一致\n", len(docMap))
		return
	}
	fmt.Printf("✗ %d/%d 个结构体有漂移\n", drift, len(docMap))
	if !*flagWrite {
		fmt.Println("  加 -write 可按文档补齐（类型映射与 tools/gen_struct.go 一致）")
	}
	os.Exit(1)
}

func (r docRow) note() string {
	if r.Note == "" {
		return ""
	}
	return "。" + r.Note
}

/* ------------------------------------------------------------ Go 侧 */

var structRe = regexp.MustCompile(`(?m)^type (\w+) struct \{`)
var tagRe = regexp.MustCompile("json:\"([^\",]+)")

func loadStructs(root string) (map[string]*goStruct, error) {
	files, err := filepath.Glob(filepath.Join(root, "*.go"))
	if err != nil {
		return nil, err
	}

	out := map[string]*goStruct{}
	fset := token.NewFileSet()

	for _, file := range files {
		if strings.HasSuffix(file, "_test.go") {
			continue
		}
		src, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}
		// 用 go/parser 拿行号；字段本身用正则取，因为标签里可能有逗号
		f, err := parser.ParseFile(fset, file, src, 0)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", file, err)
		}
		ast.Inspect(f, func(n ast.Node) bool {
			ts, ok := n.(*ast.TypeSpec)
			if !ok {
				return true
			}
			if _, ok := ts.Type.(*ast.StructType); !ok {
				return true
			}
			gs := &goStruct{
				Name: ts.Name.Name,
				File: filepath.Base(file),
				Line: fset.Position(ts.Pos()).Line,
				Tags: map[string]bool{},
			}
			// 结构体正文：从名字那行到匹配的右花括号
			body := structBody(string(src), ts.Name.Name)
			for _, m := range tagRe.FindAllStringSubmatch(body, -1) {
				if gs.Tags[m[1]] {
					continue
				}
				gs.Tags[m[1]] = true
				gs.Fields = append(gs.Fields, m[1])
			}
			out[gs.Name] = gs
			return true
		})
	}
	return out, nil
}

// structBody 取出 `type <name> struct {` 与配对的 `}` 之间那段文本。
func structBody(src, name string) string {
	head := "type " + name + " struct {"
	i := strings.Index(src, head)
	if i < 0 {
		return ""
	}
	depth := 0
	for j := i + len(head) - 1; j < len(src); j++ {
		switch src[j] {
		case '{':
			depth++
		case '}':
			depth--
			if depth == 0 {
				return src[i : j+1]
			}
		}
	}
	return ""
}

/* ---------------------------------------------------------- 文档侧 */

// matchLabel 在最近两行里找包含 want 的那一行，返回该行原文（用于诊断输出）
func matchLabel(recent [2]string, want string) (string, bool) {
	for i := len(recent) - 1; i >= 0; i-- {
		if recent[i] != "" && strings.Contains(recent[i], want) {
			return recent[i], true
		}
	}
	return "", false
}

// readTables 读多处锚点并取并集。
//
// `Label`（type.go）被多个文件多处共用，只读一处的话其余几处的字段会被
// 报成「Go 有、文档没有」。并集是有意为之：共享结构体的字段集合就是各处的并集。
//
// ⚠️ 一处锚点读不到**不一定**是写错 —— 共享结构体挂多处时，未必每处文档
// 都把它单列成一张表（比如 `vip` 对象在有的接口下是 `staff[]` 里的子对象，
// 没有独立表）。所以不致命，但**要报出来**：失效的锚点等于你以为在盯着
// 的东西其实没盯，而它的症状是「什么都没发生」，跟正常一模一样。
func readTables(docsRoot string, locs []docLoc) (merged *docTable, failed []string, allFailed error) {
	merged = &docTable{Rows: map[string]docRow{}}

	for _, loc := range locs {
		t, err := readTable(docsRoot, loc)
		if err != nil {
			failed = append(failed, fmt.Sprintf("%s ## %s → %s：%v", loc.File, loc.Section, loc.Label, err))
			continue
		}
		merged.Labels = append(merged.Labels, t.Labels...)
		for _, k := range t.Order {
			if _, dup := merged.Rows[k]; dup {
				continue
			}
			merged.Rows[k] = t.Rows[k]
			merged.Order = append(merged.Order, k)
		}
	}

	if len(merged.Order) == 0 {
		return nil, failed, fmt.Errorf("%s", strings.Join(failed, "；"))
	}
	return merged, failed, nil
}

// readTable 在文档里找到 loc 指定的那张字段表。
//
// 定位分两步：先按 `## Section` 切出章节，再在章节里找**前一行标签包含 loc.Label**
// 的表。标签用包含匹配是刻意的 —— 文档里同一个对象在不同章节写法会有小差异
// （多个空格、多一层反引号），写死全等会脆。
func readTable(docsRoot string, loc docLoc) (*docTable, error) {
	path := filepath.Join(docsRoot, filepath.FromSlash(loc.File))
	raw, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("读不到 %s（-docs 指对了吗）", loc.File)
	}

	lines := strings.Split(strings.ReplaceAll(string(raw), "\r\n", "\n"), "\n")

	// 1. 切章节
	start, end := -1, len(lines)
	for i, l := range lines {
		if !strings.HasPrefix(l, "## ") {
			continue
		}
		if start >= 0 {
			end = i
			break
		}
		if strings.TrimSpace(strings.TrimPrefix(l, "## ")) == loc.Section {
			start = i
		}
	}
	if start < 0 {
		return nil, fmt.Errorf("文档里找不到章节「%s」", loc.Section)
	}

	// 2. 在章节里找标签匹配的表
	//
	// 留**两行**窗口，不是一行。文档的表头常写成「标签行 + 说明行」两行：
	//
	//	`data`中的`Related`数组中的对象：
	//	基本同「获取视频详细信息(web端)」中的data对象，…另有以下字段：
	//	| 字段 | 类型 | … |
	//
	// 只记最近一行的话，标签行会被说明行顶掉，于是这张表匹不中，
	// 循环继续往后走、匹到 `ai_rcmd` 那张子表上 —— 而它恰好也叫
	// 「`Related`数组中的对象中的`ai_rcmd`对象」，包含匹配下就会中招。
	var recent [2]string
	for i := start; i < end; i++ {
		l := strings.TrimSpace(lines[i])

		if !strings.HasPrefix(l, "|") {
			if l == "" || strings.HasPrefix(l, "#") || strings.HasPrefix(l, "```") {
				continue
			}
			recent[0], recent[1] = recent[1], l
			continue
		}

		label, hit := matchLabel(recent, loc.Label)
		if !hit {
			// 表本身就是一道分隔：这张表之前的标签不再适用于它之后的表
			recent[0], recent[1] = "", ""
			continue
		}

		t := &docTable{Labels: []string{label}, Rows: map[string]docRow{}}
		for ; i < end; i++ {
			row := strings.TrimSpace(lines[i])
			if !strings.HasPrefix(row, "|") {
				break
			}
			cells := strings.Split(strings.Trim(row, "|"), "|")
			if len(cells) < 4 {
				continue
			}
			for k := range cells {
				cells[k] = strings.TrimSpace(strings.ReplaceAll(cells[k], "`", ""))
			}
			// 类型列可能是空的（`| premiere | | null | |`），那不是废弃行，
			// 只是文档作者没填类型。只有类型和内容**都**空才算垃圾行
			if strings.HasPrefix(cells[1], "---") || cells[1] == "类型" {
				continue
			}
			if cells[1] == "" && cells[2] == "" {
				continue
			}
			// 类型列不设白名单。语料里有 `web 端：null<br />APP 端：array`、
			// `禁用时：null<br />正常时：array`、`有效时：obj`、`null` 这些写法，
			// 白名单漏一种就静默少一行 —— 而少一行意味着这个字段**永远不会**
			// 被报「文档有、Go 没有」。表头与分隔行上面已经挡掉了，
			// 再往下能走到这里的都是正文行。
			// 字段名可能带后缀说明，如 `attribute`(已删除)
			name := cells[0]
			// 文档用删除线标退休字段（`~~attribute~~(已经弃用)`），不补进代码
			if strings.HasPrefix(name, "~~") {
				continue
			}
			if sp := strings.IndexAny(name, "(（ "); sp > 0 {
				name = strings.TrimSpace(name[:sp])
			}
			if name == "" || !fieldNameRe.MatchString(name) {
				continue
			}
			t.Rows[name] = docRow{Kind: cells[1], Content: cells[2], Note: cells[3]}
			t.Order = append(t.Order, name)
		}
		return t, nil
	}

	return nil, fmt.Errorf("章节「%s」里找不到标签含「%s」的表", loc.Section, loc.Label)
}

// fieldNameRe 匹配字段名单元格
var fieldNameRe = regexp.MustCompile(`^[A-Za-z_0-9][A-Za-z0-9_]*$`)

/* ------------------------------------------------------------ 比对 */

// knownSet 把 docKnown 的分组摊成字段 → 原因 的查找表
func knownSet(groups []docKnownGroup) map[string]string {
	if len(groups) == 0 {
		return nil
	}
	out := map[string]string{}
	for _, g := range groups {
		for _, f := range g.Fields {
			out[f] = g.Why
		}
	}
	return out
}

// diff 比对结构体与文档表。
//
// known 是 docKnown 里记过的字段（文档表没有、但已核对确实存在），
// 从 extra 里扣掉，并把实际扣掉的键回传，好让调用方报个数。
func diff(gs *goStruct, t *docTable, known map[string]string) (missing, extra, suppressed []string) {
	for _, k := range t.Order {
		if !gs.Tags[k] {
			missing = append(missing, k)
		}
	}
	for _, k := range gs.Fields {
		if _, ok := t.Rows[k]; ok {
			continue
		}
		if _, ok := known[k]; ok {
			suppressed = append(suppressed, k)
			continue
		}
		extra = append(extra, k)
	}
	return missing, extra, suppressed
}

/* ------------------------------------------------------------ 写入 */

// writeFields 把缺失字段插到结构体右花括号前，再用 go/format 格式化整个文件。
func writeFields(root string, gs *goStruct, missing []string, t *docTable) error {
	path := filepath.Join(root, gs.File)
	src, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	body := structBody(string(src), gs.Name)
	if body == "" {
		return fmt.Errorf("找不到 %s 的正文", gs.Name)
	}

	var b strings.Builder
	for _, k := range missing {
		r := t.Rows[k]
		b.WriteString("\t")
		b.WriteString(snakeToCamel(k))
		b.WriteString(" ")
		b.WriteString(goType(k, r.Kind))
		b.WriteString(" `json:\"")
		b.WriteString(k)
		b.WriteString("\"` // ")
		b.WriteString(firstLine(r.Content + r.note()))
		b.WriteString("\n")
	}

	// 插在最后一个 `}` 前
	idx := strings.LastIndex(body, "}")
	patched := strings.Replace(string(src), body, body[:idx]+b.String()+body[idx:], 1)

	formatted, err := format.Source([]byte(patched))
	if err != nil {
		return fmt.Errorf("生成后格式化失败（结构体里可能有语法问题）：%w", err)
	}
	return os.WriteFile(path, formatted, 0o644)
}

// goType 按类型单元格推一个 Go 类型。
//
// 判据是**子串包含**而不是全等：语料里同一列的写法有 `array`、`array(obj)`、
// `有效时：obj<br />无效时：null`、`禁用时：null<br />正常时：array`、
// `web 端：null<br />APP 端：array`，写死全等会一路掉到 any。
// 拿不准的（`null`、纯说明文字）仍然落到 any —— 由人看一眼再定。
func goType(jsonName, kind string) string {
	if m := arrayOfRe.FindStringSubmatch(kind); m != nil {
		switch m[1] {
		case "num", "int":
			return "[]int"
		case "str":
			return "[]string"
		case "bool":
			return "[]bool"
		case "float":
			return "[]float64"
		}
		return "[]" + singular(snakeToCamel(jsonName))
	}

	switch {
	case strings.Contains(kind, "array"):
		return "[]" + singular(snakeToCamel(jsonName))
	case strings.Contains(kind, "num"), strings.Contains(kind, "int"):
		return "int"
	case strings.Contains(kind, "bool"):
		return "bool"
	case strings.Contains(kind, "str"):
		return "string"
	case strings.Contains(kind, "float"):
		return "float64"
	case strings.Contains(kind, "obj"):
		return snakeToCamel(jsonName)
	}
	return "any"
}

var arrayOfRe = regexp.MustCompile(`^array\((\w+)\)$`)

// singular 把字段名变单数，给数组元素命名。`replies` → `Reply`，不是 `Replie`
func singular(s string) string {
	switch {
	case strings.HasSuffix(s, "ies") && len(s) > 3:
		return s[:len(s)-3] + "y"
	case strings.HasSuffix(s, "ses"), strings.HasSuffix(s, "hes"):
		return s[:len(s)-2]
	case strings.HasSuffix(s, "s") && !strings.HasSuffix(s, "ss") && len(s) > 1:
		return s[:len(s)-1]
	}
	return s
}

func snakeToCamel(s string) string {
	parts := strings.FieldsFunc(s, func(r rune) bool { return r == '_' || r == '.' || r == '-' })
	for i, w := range parts {
		if w == "" {
			continue
		}
		parts[i] = strings.ToUpper(w[:1]) + w[1:]
	}
	return strings.Join(parts, "")
}

/* ------------------------------------------------------------ 杂项 */

func reportCoverage(root string, structs map[string]*goStruct) {
	inMap := map[string]bool{}
	for k := range docMap {
		inMap[k] = true
	}

	var unmapped []string
	for name := range structs {
		if !inMap[name] && len(structs[name].Fields) > 0 {
			unmapped = append(unmapped, name)
		}
	}
	sort.Strings(unmapped)

	fmt.Printf("结构体 %d 个，已挂文档锚点 %d 个，未挂 %d 个\n\n",
		len(structs), len(docMap), len(unmapped))
	fmt.Println("未挂锚点的（挂上之后 -check 才能盯住它们的漂移）：")
	for _, n := range unmapped {
		fmt.Printf("  %-34s %s:%d\n", n, structs[n].File, structs[n].Line)
	}
}

func firstLine(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	s = strings.ReplaceAll(s, "<br />", "。")
	s = strings.TrimSpace(s)
	if r := []rune(s); len(r) > 60 {
		return string(r[:60]) + "…"
	}
	return s
}

func sortedKeys[V any](m map[string]V) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

func must(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "✗", err)
		os.Exit(2)
	}
}

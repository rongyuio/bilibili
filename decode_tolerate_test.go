package bilibili

import (
	"bytes"
	"encoding/json"
	"reflect"
	"slices"
	"testing"
)

type tolerateInner struct {
	Value int `json:"value"`
}

type tolerateTarget struct {
	Text    string          `json:"text"`
	Number  json.Number     `json:"number"`
	Quoted  int             `json:"quoted,string"`
	Flag    bool            `json:"flag"`
	Pointer *int            `json:"pointer"`
	Struct  tolerateInner   `json:"struct"`
	Slice   []int           `json:"slice"`
	Array   [2]int          `json:"array"`
	Bytes   []byte          `json:"bytes"`
	Any     any             `json:"any"`
	Raw     json.RawMessage `json:"raw"`
	Loose   NumberOrString  `json:"loose"`
	Strings map[string]int  `json:"strings"`
	IntKeys map[int]int     `json:"int_keys"`
}

type tolerateCase struct {
	name string
	in   string
	want []string
	// skipDiagnose 标记容错与诊断刻意不一致的用例，见表内说明。
	skipDiagnose bool
}

func tolerateCases() []tolerateCase {
	return []tolerateCase{
		{name: "字符串字段收到数字", in: `{"text":1}`, want: []string{"$.data.text"}},
		{name: "json.Number 收到布尔", in: `{"number":true}`, want: []string{"$.data.number"}},
		{name: "带 string 标签收到非字符串", in: `{"quoted":5}`, want: []string{"$.data.quoted"}},
		{name: "布尔字段收到数字", in: `{"flag":1}`, want: []string{"$.data.flag"}},
		{name: "指针字段收到布尔", in: `{"pointer":true}`, want: []string{"$.data.pointer"}},
		{name: "结构体字段收到数组", in: `{"struct":[1,2]}`, want: []string{"$.data.struct"}},
		{name: "结构体内层漂移", in: `{"struct":{"value":true}}`, want: []string{"$.data.struct.value"}},
		{name: "切片字段收到对象", in: `{"slice":{}}`, want: []string{"$.data.slice"}},
		{name: "切片元素漂移", in: `{"slice":[1,"x",3]}`, want: []string{"$.data.slice[1]"}},
		{name: "数组元素漂移", in: `{"array":[1,true]}`, want: []string{"$.data.array[1]"}},
		{name: "字节切片收到布尔", in: `{"bytes":true}`, want: []string{"$.data.bytes"}},
		{name: "map 值漂移", in: `{"strings":{"a":"x"}}`, want: []string{"$.data.strings.a"}},
		{
			name: "一处响应多处漂移",
			in:   `{"text":1,"slice":[true],"struct":{"value":"x"}}`,
			want: []string{"$.data.text", "$.data.slice[0]", "$.data.struct.value"},
		},
		{name: "null 一律可接受", in: `{"pointer":null,"slice":null,"number":null,"loose":null}`},
		{name: "any 与 RawMessage 不设限", in: `{"any":{"deep":[1,{"x":true}]},"raw":{"x":true}}`},
		{name: "NumberOrString 接受布尔", in: `{"loose":true}`},
		{name: "未知键被忽略", in: `{"unknown":true}`},
		{
			// 刻意分歧：map 键转换失败要靠删成员才能修，本层不实现，诊断仍然精确定位。
			name:         "map 键无法转换时容错不介入",
			in:           `{"int_keys":{"abc":1}}`,
			skipDiagnose: true,
		},
	}
}

func TestTolerateDecodePruneRules(t *testing.T) {
	target := reflect.TypeFor[tolerateTarget]()
	for _, c := range tolerateCases() {
		t.Run(c.name, func(t *testing.T) {
			pruned, dropped, changed := tolerateDecode([]byte(c.in), target, 0, true, true)
			wantChanged := len(c.want) > 0
			if changed != wantChanged {
				t.Fatalf("changed = %t, want %t", changed, wantChanged)
			}
			if paths := droppedPaths(dropped); !slices.Equal(paths, c.want) {
				t.Errorf("丢弃路径 = %v, want %v", paths, c.want)
			}
			if !changed {
				return
			}
			// 剪枝结果必须真的能解开，否则容错机制本身就不成立。
			if err := json.Unmarshal(pruned, new(tolerateTarget)); err != nil {
				t.Errorf("剪枝后的字节仍无法解码: %v\n%s", err, pruned)
			}
		})
	}
}

// TestPruneAndDiagnoseAgree 保证「诊断认为能精确定位」与「容错会剪掉」结论一致，
// 两层各自维护一份遍历逻辑时，这是防止它们漂移的唯一护栏。
func TestPruneAndDiagnoseAgree(t *testing.T) {
	target := reflect.TypeFor[tolerateTarget]()
	for _, c := range tolerateCases() {
		if c.skipDiagnose {
			continue
		}
		t.Run(c.name, func(t *testing.T) {
			_, dropped, _ := tolerateDecode([]byte(c.in), target, 0, true, true)
			location, exact := diagnoseValue([]byte(c.in), target, "$.data", "", 0, false)
			switch {
			case len(dropped) == 0 && location == nil:
			case len(dropped) == 0:
				t.Errorf("诊断报出 %s，容错却没有剪枝", location.JSONPath)
			case !exact:
				t.Errorf("容错剪掉 %v，诊断却只能定位到边界", droppedPaths(dropped))
			case location.JSONPath != dropped[0].JSONPath:
				t.Errorf("首个丢弃 %s 与诊断位置 %s 不一致", dropped[0].JSONPath, location.JSONPath)
			}
		})
	}
}

// TestDecodeNullIntoNonNullableIgnored 钉住整套剪枝依赖的 encoding/json 语义：
// null 对非指针标量、结构体、数组不报错也不覆盖，对指针/切片/map 置 nil。
func TestDecodeNullIntoNonNullableIgnored(t *testing.T) {
	const allNull = `{"text":null,"number":null,"quoted":null,"flag":null,"pointer":null,` +
		`"struct":null,"slice":null,"array":null,"bytes":null,"any":null,"raw":null,` +
		`"loose":null,"strings":null,"int_keys":null}`

	var fresh tolerateTarget
	if err := json.Unmarshal([]byte(allNull), &fresh); err != nil {
		t.Fatalf("null 不应导致解码失败: %v", err)
	}
	// json.RawMessage 保留原始字节 "null"，其余字段均为零值。
	if !reflect.DeepEqual(fresh, tolerateTarget{Raw: json.RawMessage(nullLiteral)}) {
		t.Errorf("null 应保持零值: %+v", fresh)
	}

	// 覆盖语义的不对称正是「重试必须解码到全新值」的原因。
	number := 5
	prefilled := tolerateTarget{
		Text:    "keep",
		Number:  json.Number("7"),
		Quoted:  9,
		Flag:    true,
		Pointer: &number,
		Slice:   []int{1, 2},
	}
	if err := json.Unmarshal([]byte(allNull), &prefilled); err != nil {
		t.Fatalf("null 不应导致解码失败: %v", err)
	}
	if prefilled.Text != "keep" || prefilled.Number.String() != "7" || prefilled.Quoted != 9 || !prefilled.Flag {
		t.Errorf("null 不应覆盖标量字段: %+v", prefilled)
	}
	if prefilled.Pointer != nil || prefilled.Slice != nil {
		t.Errorf("null 应把指针与切片置 nil: %+v", prefilled)
	}
}

func TestTolerateDecodeKeepsDuplicateKeys(t *testing.T) {
	// 同一个键出现两次，只有后一次类型错误；剪枝后前一次仍然生效。
	data := []byte(`{"text":"a","text":1}`)
	pruned, dropped, changed := tolerateDecode(data, reflect.TypeFor[tolerateTarget](), 0, true, true)
	if !changed || len(dropped) != 1 {
		t.Fatalf("changed = %t, dropped = %v", changed, dropped)
	}
	var target tolerateTarget
	if err := json.Unmarshal(pruned, &target); err != nil {
		t.Fatalf("剪枝后解码失败: %v", err)
	}
	if target.Text != "a" {
		t.Errorf("Text = %q, want a", target.Text)
	}
}

func TestTolerateDecodeNoChangeSkipsRetry(t *testing.T) {
	data := []byte(`{"text":"a","number":1}`)
	pruned, dropped, changed := tolerateDecode(data, reflect.TypeFor[tolerateTarget](), 0, true, true)
	if changed || dropped != nil {
		t.Fatalf("changed = %t, dropped = %v; want false, nil", changed, dropped)
	}
	if &pruned[0] != &data[0] {
		t.Error("无漂移时不应重建字节")
	}
}

func TestTolerateDecodeRootNeverPrunes(t *testing.T) {
	data := []byte(`[]`)
	pruned, dropped, changed := tolerateDecode(data, reflect.TypeFor[tolerateTarget](), 0, true, true)
	if changed || len(dropped) != 0 {
		t.Fatalf("根节点不应剪枝: changed = %t, dropped = %v", changed, dropped)
	}
	if &pruned[0] != &data[0] {
		t.Error("根节点不剪枝时应保留原字节")
	}
}

func TestTolerateDecodeOffset(t *testing.T) {
	target := reflect.TypeFor[tolerateTarget]()
	data := []byte(`{"text":1}`)

	_, dropped, _ := tolerateDecode(data, target, 0, false, true)
	if len(dropped) != 1 || dropped[0].Offset != 0 {
		t.Errorf("data 偏移不可用时 Offset 应为 0: %+v", dropped)
	}

	const base = 200
	_, dropped, _ = tolerateDecode(data, target, base, true, true)
	if want := base + int64(bytes.Index(data, []byte("1"))) + 1; len(dropped) != 1 || dropped[0].Offset != want {
		t.Errorf("Offset = %+v, want %d", dropped, want)
	}
}

func TestTolerateDecodeWithoutReport(t *testing.T) {
	// report 为 false 仍然剪枝，只是不产生上报。
	pruned, dropped, changed := tolerateDecode([]byte(`{"text":1}`), reflect.TypeFor[tolerateTarget](), 0, true, false)
	if !changed || dropped != nil {
		t.Fatalf("changed = %t, dropped = %v; want true, nil", changed, dropped)
	}
	if err := json.Unmarshal(pruned, new(tolerateTarget)); err != nil {
		t.Errorf("剪枝后解码失败: %v", err)
	}
}

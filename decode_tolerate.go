package bilibili

import (
	"encoding/json"
	"reflect"
)

// nullLiteral 是剪枝时用来替换不兼容节点的字节。
const nullLiteral = "null"

// DroppedField 描述一个因 JSON 类型与 Go 类型不兼容而被丢弃的字段。
// 它对应一次成功但降级的解码结果，不是错误；不含字段值，可以安全写入日志。
type DroppedField struct {
	Method   string // HTTP 方法
	Endpoint string // 已清洗的接口地址，不含查询参数、用户信息与片段
	RootType string // 结果类型名，例如 bilibili.DynamicInfo
	GoField  string // Go 字段路径；与 DecodeError 一样不含切片下标
	JSONPath string // 例如 $.data.items[0].modules.module_author.following
	Expected string // 期望的 Go 类型名
	Actual   string // 实际的 JSON kind：number/string/boolean/object/array
	Offset   int64  // 原始响应体中从 1 开始计数的字节偏移；0 表示不可用
}

// DroppedFieldHandler 在容错解码丢弃字段时被调用。
// 同一条响应内按 JSON 文档顺序可能触发多次；不同请求之间可能并发。
// 回调在请求的 goroutine 中执行，必须并发安全。本库不捕获回调的 panic。
type DroppedFieldHandler func(DroppedField)

// SetDroppedFieldHandler 设置容错解码回调，传 nil 只关闭上报，容错仍然开启。
// 与 Cookie、会话切换等配置一样，必须在请求之外调用。
func (c *Client) SetDroppedFieldHandler(handler DroppedFieldHandler) {
	c.dropMu.Lock()
	defer c.dropMu.Unlock()
	c.dropHandler = handler
}

// droppedFieldHandler 取一次回调快照；调用方在锁外触发回调，避免回调重入死锁。
func (c *Client) droppedFieldHandler() DroppedFieldHandler {
	c.dropMu.RLock()
	defer c.dropMu.RUnlock()
	return c.dropHandler
}

// tolerateDecode 把 data 中 JSON kind 与 target 不兼容的叶子替换为 null。
// 根节点绝不剪枝：整个 data 被置零却报成功会掩盖接口形态变化。
// base 是 data 值在原始响应体中的 0 起始偏移，found 为 false 时 Offset 记为 0。
// report 为 false 时仍然剪枝，只是不构造 DroppedField。
func tolerateDecode(data []byte, target reflect.Type, base int64, found, report bool) ([]byte, []DroppedField, bool) {
	w := tolerateWalk{report: report, found: found}
	pruned, changed := w.value(data, target, false, base, "$.data", "", true)
	if !changed {
		return data, nil, false
	}
	return pruned, w.dropped, true
}

type tolerateWalk struct {
	report  bool
	found   bool
	dropped []DroppedField
}

// value 遍历一个节点，不兼容时返回 null。root 为真时只观察、不剪枝。
// 子节点未改动时丢弃本级记录，避免放弃剪枝后残留无效的上报。
func (w *tolerateWalk) value(raw []byte, t reflect.Type, quoted bool, offset int64, path, field string, root bool) ([]byte, bool) {
	mark := len(w.dropped)
	out, changed := w.walk(raw, t, quoted, offset, path, field, root)
	if !changed {
		w.dropped = w.dropped[:mark]
	}
	return out, changed
}

func (w *tolerateWalk) walk(raw []byte, t reflect.Type, quoted bool, offset int64, path, field string, root bool) ([]byte, bool) {
	// 与 decode_diagnostic.go 的 diagnoseValue 保持同序判定，两层结论必须一致。
	expected := diagnosticTypeName(t)
	for t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	if t == rawMessageType || hasDecoder(t) {
		return raw, false
	}
	kind := jsonKind(raw)
	if t == numberOrStringType {
		if numberOrStringAcceptsKind(kind) {
			return raw, false
		}
		return w.drop(expected, kind, offset, path, field, root)
	}
	if kind == "null" {
		return raw, false
	}
	if scalarTarget(t, quoted) {
		if t.Kind() == reflect.Interface || probeDecode(raw, t, quoted) == nil {
			return raw, false
		}
		return w.drop(expected, kind, offset, path, field, root)
	}
	if ((t.Kind() == reflect.Struct || t.Kind() == reflect.Map) && kind != "object") ||
		((t.Kind() == reflect.Array || t.Kind() == reflect.Slice) && kind != "array") {
		return w.drop(expected, kind, offset, path, field, root)
	}
	children := jsonChildren(raw)
	if len(children) == 0 {
		return raw, false
	}
	var fields []decodeField
	if t.Kind() == reflect.Struct {
		fields = decodeFields(t)
	}
	rebuilt := make([][]byte, 0, len(children))
	changed := false
	for i, child := range children {
		childType, childField, childPath, childQuoted := t, field, path, false
		switch t.Kind() {
		case reflect.Struct:
			f, ok := matchDecodeField(fields, child.key)
			if !ok {
				// 未知键会被 encoding/json 忽略，永远不会失败，保持原样。
				rebuilt = append(rebuilt, child.raw)
				continue
			}
			childType, childQuoted = f.typ, f.quoted
			childField = f.goPath
			if field != "" {
				childField = field + "." + childField
			}
			childPath = appendJSONKey(path, child.key)
		case reflect.Map:
			// 键转换失败要靠删成员才能修，本层不实现，保持致命。
			if hasDecoder(t.Key()) || mapKeyMismatch(t.Key(), child.key) {
				rebuilt = append(rebuilt, child.raw)
				continue
			}
			childType, childPath = t.Elem(), appendJSONKey(path, child.key)
		case reflect.Array, reflect.Slice:
			if t.Kind() == reflect.Array && i >= t.Len() {
				rebuilt = append(rebuilt, child.raw)
				continue
			}
			childType, childPath = t.Elem(), path+"["+child.key+"]"
		}
		childRaw, childChanged := w.value(child.raw, childType, childQuoted, offset+child.offset, childPath, childField, false)
		if childChanged {
			changed = true
		}
		rebuilt = append(rebuilt, childRaw)
	}
	if !changed {
		return raw, false
	}
	pruned := rebuildJSON(kind, children, rebuilt)
	if !json.Valid(pruned) {
		return raw, false
	}
	return pruned, true
}

// drop 记录被丢弃的字段并返回替换值；根节点始终保留原字节。
func (w *tolerateWalk) drop(expected, kind string, offset int64, path, field string, root bool) ([]byte, bool) {
	if root {
		return nil, false
	}
	if w.report {
		at := int64(0)
		if w.found {
			at = offset + 1
		}
		w.dropped = append(w.dropped, DroppedField{
			GoField:  field,
			JSONPath: path,
			Expected: expected,
			Actual:   kind,
			Offset:   at,
		})
	}
	return []byte(nullLiteral), true
}

// rebuildJSON 按原文档顺序重建对象或数组，重复键逐条保留。
func rebuildJSON(kind string, children []jsonChild, values [][]byte) []byte {
	object := kind == "object"
	pruned := make([]byte, 0, len(children)*8)
	if object {
		pruned = append(pruned, '{')
	} else {
		pruned = append(pruned, '[')
	}
	for i, child := range children {
		if i > 0 {
			pruned = append(pruned, ',')
		}
		if object {
			key, err := json.Marshal(child.key)
			if err != nil {
				return nil
			}
			pruned = append(append(pruned, key...), ':')
		}
		pruned = append(pruned, values[i]...)
	}
	if object {
		return append(pruned, '}')
	}
	return append(pruned, ']')
}

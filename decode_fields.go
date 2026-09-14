package bilibili

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type decodeField struct {
	name   string
	goPath string
	typ    reflect.Type
	index  []int
	tagged bool
	quoted bool
}

// Match encoding/json's promoted-field dominance, then its exact/folded lookup.
func decodeFields(root reflect.Type) []decodeField {
	var candidates []decodeField
	var visit func(reflect.Type, string, []int, map[reflect.Type]bool)
	visit = func(t reflect.Type, prefix string, index []int, parents map[reflect.Type]bool) {
		if parents[t] {
			return
		}
		parents[t] = true
		defer delete(parents, t)
		for i := range t.NumField() {
			f := t.Field(i)
			ft := f.Type
			if ft.Kind() == reflect.Pointer {
				ft = ft.Elem()
			}
			if !f.IsExported() && (!f.Anonymous || ft.Kind() != reflect.Struct) {
				continue
			}
			tag := f.Tag.Get("json")
			if tag == "-" {
				continue
			}
			parts := strings.Split(tag, ",")
			name := parts[0]
			if !validJSONTag(name) {
				name = ""
			}
			idx := append(append([]int(nil), index...), i)
			goPath := prefix + f.Name
			if name == "" && f.Anonymous && ft.Kind() == reflect.Struct {
				visit(ft, goPath+".", idx, parents)
				continue
			}
			tagged := name != ""
			if name == "" {
				name = f.Name
			}
			quoted := false
			for _, option := range parts[1:] {
				if option == "string" {
					quoted = true
				}
			}
			switch ft.Kind() {
			case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.String:
			default:
				quoted = false
			}
			candidates = append(candidates, decodeField{name: name, goPath: goPath, typ: f.Type, index: idx, tagged: tagged, quoted: quoted})
		}
	}
	visit(root, "", nil, make(map[reflect.Type]bool))
	groups := make(map[string][]decodeField)
	for _, f := range candidates {
		groups[f.name] = append(groups[f.name], f)
	}
	var fields []decodeField
	for _, group := range groups {
		sort.SliceStable(group, func(i, j int) bool {
			if len(group[i].index) != len(group[j].index) {
				return len(group[i].index) < len(group[j].index)
			}
			return group[i].tagged && !group[j].tagged
		})
		if len(group) > 1 && len(group[0].index) == len(group[1].index) && group[0].tagged == group[1].tagged {
			continue
		}
		fields = append(fields, group[0])
	}
	sort.Slice(fields, func(i, j int) bool {
		for k := 0; k < len(fields[i].index) && k < len(fields[j].index); k++ {
			if fields[i].index[k] != fields[j].index[k] {
				return fields[i].index[k] < fields[j].index[k]
			}
		}
		return len(fields[i].index) < len(fields[j].index)
	})
	return fields
}

func matchDecodeField(fields []decodeField, key string) (decodeField, bool) {
	for _, f := range fields {
		if f.name == key {
			return f, true
		}
	}
	for _, f := range fields {
		if strings.EqualFold(f.name, key) {
			return f, true
		}
	}
	return decodeField{}, false
}

func validJSONTag(tag string) bool {
	if tag == "" {
		return false
	}
	for _, r := range tag {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) && !strings.ContainsRune("!#$%&()*+-./:;<=>?@[]^_{|}~ ", r) {
			return false
		}
	}
	return true
}

// Keep named types recognizable without expanding anonymous response schemas.
// JSONPath and GoField carry the field-level location separately.
func diagnosticTypeName(t reflect.Type) string {
	if t.Name() != "" {
		name := t.String()
		// Generic arguments may themselves contain an entire anonymous struct.
		if strings.Contains(name, "struct {") {
			if index := strings.IndexByte(name, '['); index >= 0 {
				return name[:index] + "[...]"
			}
		}
		return name
	}
	switch t.Kind() {
	case reflect.Pointer:
		return "*" + diagnosticTypeName(t.Elem())
	case reflect.Slice:
		return "[]" + diagnosticTypeName(t.Elem())
	case reflect.Array:
		return "[" + strconv.Itoa(t.Len()) + "]" + diagnosticTypeName(t.Elem())
	case reflect.Map:
		return "map[" + diagnosticTypeName(t.Key()) + "]" + diagnosticTypeName(t.Elem())
	case reflect.Struct:
		return "struct{...}"
	case reflect.Interface:
		return "interface{...}"
	default:
		return t.Kind().String()
	}
}

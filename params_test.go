package bilibili

import "testing"

func TestToSnakeCase(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		// 初始缩写整体成词（第四轮修复的回归用例）
		{"IDs", "ids"},
		{"IDsA", "ids_a"},
		{"UIDType", "uid_type"},
		{"DynamicID", "dynamic_id"},
		{"BaseURL", "base_url"},
		{"RoomID", "room_id"},
		{"URL", "url"},
		{"URLs", "urls"},
		{"URI", "uri"},
		{"URIs", "uris"},
		{"JSON", "json"},
		{"HTTP", "http"},
		{"API", "api"},
		{"APIs", "apis"},
		{"UUID", "uuid"},
		{"UUIDs", "uuids"},
		{"ID", "id"},
		{"IDCard", "id_card"},
		// 普通单词
		{"NickName", "nick_name"},
		{"HostMid", "host_mid"},
		{"DisplayRank", "display_rank"},
		{"TimeZoneOffset", "time_zone_offset"},
		// 不要把普通单词误拆成缩写
		{"Initialization", "initialization"},
		{"Identity", "identity"},
		{"Idempotent", "idempotent"},
		// 单个词
		{"Name", "name"},
	}
	for _, c := range cases {
		if got := toSnakeCase(c.in); got != c.want {
			t.Errorf("toSnakeCase(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}

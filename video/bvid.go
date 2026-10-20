package video

import "strings"

// bv2Av 将 bv 号转换为 av 号，传入的 bv 号格式为 "BV1xxxxxxxxx"，前面的 "BV" 不区分大小写。
// 调用方必须先校验长度为 12，否则会 panic。
func bv2Av(bvid string) int {
	if len(bvid) != 12 {
		panic("bvid 格式错误: " + bvid)
	}
	const (
		xorCode  = 0x1552356C4CDB
		maskCode = 1<<51 - 1
		alphabet = "FcwAPNKTMug3GV5Lj7EJnHpWsx4tb8haYeviqBz6rkCy12mUSDQX9RdoZf"
	)
	tmp := 0
	for _, e := range []int{9, 7, 5, 6, 4, 8, 3, 10, 11} {
		idx := strings.IndexByte(alphabet, bvid[e])
		tmp = tmp*len(alphabet) + idx
	}
	return (tmp & maskCode) ^ xorCode
}

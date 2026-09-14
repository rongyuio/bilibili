package bilibili

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"slices"
)

// calculateAppSign 计算 APP API 签名
// 按照 Bilibili APP API 签名算法：参数按 key 排序后拼接，加上秘钥后计算 MD5
func calculateAppSign(params map[string]string, appSecret string) string {
	// 收集所有非空参数
	keys := make([]string, 0, len(params))
	for k, v := range params {
		if v != "" {
			keys = append(keys, k)
		}
	}

	// 按 key 排序
	slices.Sort(keys)

	// 构建查询字符串
	query := url.Values{}
	for _, k := range keys {
		if params[k] != "" {
			query.Set(k, params[k])
		}
	}

	// 拼接参数和秘钥
	signStr := query.Encode() + appSecret

	// 计算 MD5
	hash := md5.Sum([]byte(signStr))
	return hex.EncodeToString(hash[:])
}

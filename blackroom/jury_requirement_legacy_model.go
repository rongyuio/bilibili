package blackroom

// GetJuryRequirementLegacyResult 是检查申请风纪委员会资格（旧版 API）的 data 对象。
//
// 旧版比新版多一个 level 字段（会员等级 ≥ Lv4）。
type GetJuryRequirementLegacyResult struct {
	Blocked bool `json:"blocked"` // 是否被封禁。true：是。false：否
	Cert    bool `json:"cert"`    // 是否实名认证
	Level   bool `json:"level"`   // 等级是否 >= 4。仅旧版 API 返回
	Rule    bool `json:"rule"`    // 是否 90 天内无违规
}

package bilibili

// 视频分区相关响应模型。

// ZoneInfo 结构体用来表示CSV文件中的数据, 包含名称、代码、主分区tid、子分区tid,概述和备注等信息
type ZoneInfo struct {
	Name      string // 中文名称
	Code      string // 代号即英文名
	MasterTid int    // 主分区tid
	Tid       int    // 子分区tid
	Overview  string // 概述,简介
}

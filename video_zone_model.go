package bilibili

// 视频分区相关响应模型。

// ZoneInfo 表示 video_zone.csv 中的一行分区数据，主分区行与子分区行共用该结构体
// （主分区行的 MasterTid 与 Tid 相同）。
//
// CSV 的列依次为「名称、代号、主分区tid、tid、简介、备注」，前 5 列对应下列字段；
// 第 6 列「备注」在数据中始终为空，也没有对应的字段，readCSV 不会读取它。
type ZoneInfo struct {
	Name      string // 中文名称
	Code      string // 代号即英文名
	MasterTid int    // 主分区tid
	Tid       int    // 子分区tid
	Overview  string // 概述,简介
}

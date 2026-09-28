package main

// docLoc 指向文档里的一张字段表。
//
//	File    相对接口文档仓库根的 Markdown 路径
//	Section `## ` 那一行的标题原文
//	Label   表前面那行标签的**包含匹配**子串，如 "`data`中的`control`对象"
//
// Section 为什么不省略：同一个文件里 `data`对象 会出现十几次（每个接口一段），
// 只按标签找必然匹错。
type docLoc struct {
	File    string
	Section string
	Label   string
}

// docMap 结构体 → 文档位置。**一个结构体可以挂多处** —— `Label`（type.go）
// 被 4 个文件 5 处共用，只挂一处的话其余几处的字段会被报成「Go 有、文档没有」。
// 挂多处时按**并集**比对。
//
// ## 维护约定
//
// 新增接口时，给它的每个响应结构体挂一条。没挂的结构体 -check 盯不到，
// `-coverage` 会把它们列出来。
//
// ⚠️ 锚点是**人的判断**，不要照着名字猜。文档里的对象叫 `data`、`list[]`、
// `data`中的`control`对象，而 Go 结构体叫 CommentsControl、VipUserVip ——
// 名字对不上。猜错不会报错，只会把字段加到错的结构体上，然后静默上线。
// 拿不准就去看文档那一节的结构，或者留空不挂。
var docMap = map[string][]docLoc{
	/* ── 通用 ─────────────────────────────────────────────────────────── */
	"Label": {
		{File: "docs/user/info.md", Section: "用户空间详细信息", Label: "`vip`中的`label`对象"},
		// 视频这边同一张表挂在 `staff[]` 下，标签原文是
		// 「`staff`数组中的对象中的`vip`对象中的`label`对象」，
		// 写成上面的 `vip`中的`label`对象 匹不中（多了「对象」两个字）
		{File: "docs/video/info.md", Section: "获取视频详细信息(web端)", Label: "`vip`对象中的`label`对象"},
	},
	"LabelGoto": {
		{File: "docs/user/info.md", Section: "用户空间详细信息", Label: "`label_goto`对象"},
	},

	/* ── 视频 ─────────────────────────────────────────────────────────── */
	"VideoInfo": {
		{File: "docs/video/info.md", Section: "获取视频详细信息(web端)", Label: "`data`对象"},
		// 「推荐视频」的对象文档写明「基本同 data 对象，另有以下字段」，
		// 那 11 个字段是本结构体的一部分，挂第二处才不会天天报「Go 有、文档没有」
		{File: "docs/video/info.md", Section: "获取视频超详细信息(web端)", Label: "`Related`数组中的对象"},
	},
	"VideoDetailInfo": {
		{File: "docs/video/info.md", Section: "获取视频超详细信息(web端)", Label: "`data`对象"},
	},
	"VideoOnlineInfo": {
		{File: "docs/video/online.md", Section: "获取视频在线人数_web端", Label: "`data`对象"},
	},
	"CardVip": {
		{File: "docs/user/info.md", Section: "用户空间详细信息", Label: "`data`中的`vip`对象"},
		// 视频这边是 `staff[]` 里那个 vip，标签原文带「对象」两个字
		{File: "docs/video/info.md", Section: "获取视频详细信息(web端)", Label: "`vip`对象："},
	},
	"VipOttInfo": {
		{File: "docs/user/info.md", Section: "用户空间详细信息", Label: "`vip`中的`ott_info`对象"},
	},
	"VipSuperVip": {
		{File: "docs/user/info.md", Section: "用户空间详细信息", Label: "`vip`中的`super_vip`对象"},
	},
	"CollectionVideo": {
		{File: "docs/video/collection.md", Section: "获取视频合集信息", Label: "`archives`数组中的对象"},
	},

	/* ── 专栏 ─────────────────────────────────────────────────────────── */
	"Article": {
		{File: "docs/article/articles.md", Section: "获取文集基本信息", Label: "`data`中的`articles`数组中的对象"},
	},
	"ArticleInfo": {
		{File: "docs/article/info.md", Section: "获取专栏文章基本信息", Label: "`data`对象"},
	},

	/* ── 评论 ─────────────────────────────────────────────────────────── */
	"CommentsControl": {
		{File: "docs/comment/list.md", Section: "获取评论区明细_翻页加载", Label: "`data`中的`control`对象"},
	},
	"CommentsDetail": {
		{File: "docs/comment/list.md", Section: "获取评论区明细_翻页加载", Label: "`data`对象"},
	},

	/* ── 用户 ─────────────────────────────────────────────────────────── */
	"UserSpaceDetail": {
		{File: "docs/user/info.md", Section: "用户空间详细信息", Label: "`data`对象"},
	},
	"MyUserSpaceDetail": {
		{File: "docs/user/info.md", Section: "登录用户空间详细信息", Label: "`data`对象"},
	},
	"Profession": {
		{File: "docs/user/info.md", Section: "用户空间详细信息", Label: "`data`中的`profession`对象"},
	},
	"MyHonours": {
		{File: "docs/user/info.md", Section: "登录用户空间详细信息", Label: "`data`中的`honours`对象"},
	},
	"MyHonourColour": {
		{File: "docs/user/info.md", Section: "登录用户空间详细信息", Label: "`honours`中的`colour`对象"},
	},
	"MyLevelExp": {
		{File: "docs/user/info.md", Section: "登录用户空间详细信息", Label: "`data`中的`level_exp`对象"},
	},
	"MyAttestation": {
		{File: "docs/user/info.md", Section: "登录用户空间详细信息", Label: "`data`中的`attestation`对象"},
	},
	"MyAttestationCommonInfo": {
		{File: "docs/user/info.md", Section: "登录用户空间详细信息", Label: "`attestation`中的`common_info`对象"},
	},
	"MyAttestationSpliceInfo": {
		{File: "docs/user/info.md", Section: "登录用户空间详细信息", Label: "`attestation`中的`splice_info`对象"},
	},

	/* ── 收藏夹 ───────────────────────────────────────────────────────── */
	"FavourFolderInfo": {
		{File: "docs/fav/list.md", Section: "获取收藏夹内容明细列表", Label: "`data`中的`info`对象"},
	},
	"FavourMedia": {
		{File: "docs/fav/list.md", Section: "获取收藏夹内容明细列表", Label: "`medias`数组中的对象"},
	},
	"FavourResourceCntInfo": {
		{File: "docs/fav/list.md", Section: "获取收藏夹内容明细列表", Label: "`medias`数组中的对象中的`cnt_info`对象"},
	},

	/* ── 登录 ─────────────────────────────────────────────────────────── */
	"AccountInformation": {
		{File: "docs/login/member_center.md", Section: "获取我的信息", Label: "`data`对象"},
	},
	"CaptchaResult": {
		{File: "docs/login/login_action/readme.md", Section: "验证登录", Label: "`data`对象"},
	},
	"Tencent": {
		{File: "docs/login/login_action/readme.md", Section: "验证登录", Label: "`tencent`对象"},
	},

	/* ── 大会员 ───────────────────────────────────────────────────────── */
	"VipPrivilege": {
		{File: "docs/vip/info.md", Section: "卡券状态查询", Label: "`data`对象"},
	},
	"VipPrivilegeInfo": {
		{File: "docs/vip/info.md", Section: "卡券状态查询", Label: "`list`数组中的对象"},
	},
	"ExpParams": {
		{File: "docs/vip/info.md", Section: "卡券状态查询", Label: "`list[].exp_params` 对象"},
	},
	"ExtraParams": {
		{File: "docs/vip/info.md", Section: "卡券状态查询", Label: "`list[].extra_params` 对象"},
	},
	"ComicShowCoupon": {
		{File: "docs/vip/info.md", Section: "卡券状态查询", Label: "`comic_show_coupon` 对象"},
	},
	"CouponInfo": {
		{File: "docs/vip/info.md", Section: "卡券状态查询", Label: "`coupons` 数组中的对象"},
	},
	"VipUserVip": {
		{File: "docs/vip/center.md", Section: "大会员中心信息", Label: "`user`中的`vip`对象"},
	},
	"VipCenterInfo": {
		{File: "docs/vip/center.md", Section: "大会员中心信息", Label: "`data`对象"},
	},
	"FreeWelfare": {
		{File: "docs/vip/center.md", Section: "大会员中心信息", Label: "`free_welfare`数组中的对象"},
	},
	"ExtraParamas": {
		{File: "docs/vip/center.md", Section: "大会员中心信息", Label: "`data`中的`extra_paramas`对象"},
	},
}

// docKnownGroup 是一组「文档表里没有、但确认存在」的字段，外加一条为什么。
type docKnownGroup struct {
	Fields []string
	Why    string
}

// docKnown 记「Go 里有、当前锚点的表里没有」但已核对过的字段。
//
// 不并进锚点、也不删字段，是因为三种成因都不是锚点写错：
//
//	文档表漏字段     表里有它的兄弟项，就是没列它
//	字段在别处有表   结构体服务多个接口，文档只在另一个章节写了那张表
//	抓包有、文档没写 只在本库的实机响应里见过
//
// 记在这里而不是留着红灯，是因为红灯久了就没人看了 —— 新漂移会淹在已知的里面。
// 删字段前先翻这里：这些是**确认存在**的，不是残留。
var docKnown = map[string][]docKnownGroup{
	"Article": {{
		Fields: []string{
			"banner_url", "template_id", "author", "reprint", "ctime", "tags",
			"dynamic", "origin_image_urls", "is_like", "media", "apply_time",
			"check_time", "original", "act_id", "cover_avid", "type",
		},
		Why: "字段表在 docs/article/card.md 的「获取专栏显示卡片信息」一节" +
			"（`data` 中代表专栏的对象）。那张表还列了 attributes、dispute、" +
			"top_video_info、content_pic_list 等本结构体不该收的字段，并进锚点会把它们一起拉进来。" +
			"本结构体只锚了文集那条（articles[] 是专栏对象的一个子集）",
	}},
	"MyUserSpaceDetail": {
		{
			Fields: []string{"pendant", "nameplate", "official"},
			Why: "文档在本节的 `data` 表下方另开了三张子表（`data`中的`pendant`对象、" +
				"`data`中的`nameplate`对象、`data`中的`Official`对象），它们是 data 的字段，" +
				"只是没写进 data 表那一行",
		},
		{
			Fields: []string{"in_reg_audit", "profession"},
			Why:    "只写在「用户空间详细信息」（非登录版）的 `data` 表里，本结构体锚的是登录版",
		},
	},
	"UserSpaceDetail": {{
		Fields: []string{"tags", "mcn_info", "gaia_data"},
		Why:    "实机响应里有，文档的 `data` 表没列",
	}},
	"FavourMedia": {{
		Fields: []string{"ugc"},
		Why:    "实机响应里有，docs/fav/list.md 全篇没出现过这个字段",
	}},
	"VideoInfo": {{
		Fields: []string{"season_id", "premiere"},
		Why: "实机响应里有。文档只在 ugc_season 各表和 JSON 示例里出现这两个名字，" +
			"`data` 表本身没列（premiere 在表里但不是被锚点那张表收的）",
	}},
	"VideoDetailInfo": {{
		Fields: []string{"Spec", "elec", "recommend"},
		Why: "实机响应里有，文档的 `data` 表没列。注意这三个是大写开头 —— " +
			"文档的键就是 `Spec`，不是 `spec`",
	}},
}

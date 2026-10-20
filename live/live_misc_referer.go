package live

import (
	"github.com/go-resty/resty/v2"
	"github.com/rongyuio/bilibili"
)

// liveMiscReferer 给**直播域**的接口设 Referer / Origin。
//
// 直播接口走 `api.live.bilibili.com`，而真实调用发生在直播间页面上
// （`https://live.bilibili.com/<房间号>`）。与 `New()` 给的主站头
// （`https://www.bilibili.com`）对不上时，风控会回业务码 -352。
//
// 本批接口都拿不到房间号，统一退成直播首页（域对得上即可）。
// 根包的 liveReferer 未导出，子包按同样语义自行实现。
func liveMiscReferer() bilibili.Handler {
	return func(r *resty.Request) error {
		r.SetHeader("Referer", "https://live.bilibili.com/")
		r.SetHeader("Origin", "https://live.bilibili.com")
		return nil
	}
}

package bilibili

import (
	"regexp"
	"testing"

	"github.com/go-resty/resty/v2"
)

func restyNew() *resty.Request {
	return resty.New().R()
}

func TestCheckLiveAnchorLotteryParamEncoding(t *testing.T) {
	r := restyNew()
	if err := withParams(r, CheckLiveAnchorLotteryParam{RoomID: 1234567}); err != nil {
		t.Fatalf("withParams 返回错误: %v", err)
	}
	if r.QueryParam.Get("roomid") != "1234567" {
		t.Fatalf("接口参数名应为 roomid（无下划线），实际: %v", r.QueryParam)
	}
}

func TestRandomVisitID(t *testing.T) {
	visitID, err := randomVisitID()
	if err != nil {
		t.Fatalf("randomVisitID 返回错误: %v", err)
	}
	if len(visitID) != 12 {
		t.Fatalf("randomVisitID 长度应为 12，实际 %q", visitID)
	}
	if visitID[0] < '1' || visitID[0] > '9' {
		t.Fatalf("randomVisitID 首位应为 1-9，实际 %q", visitID)
	}
	matched, _ := regexp.MatchString(`^[0-9a-z]{10}$`, visitID[1:11])
	if !matched {
		t.Fatalf("randomVisitID 中间 10 位应为小写字母数字，实际 %q", visitID[1:11])
	}
	if visitID[11] != '0' {
		t.Fatalf("randomVisitID 末位应为 0，实际 %q", visitID)
	}
	again, err := randomVisitID()
	if err != nil {
		t.Fatalf("randomVisitID 第二次返回错误: %v", err)
	}
	if visitID == again {
		t.Fatalf("两次生成的 visitID 不应相同: %q", visitID)
	}
}

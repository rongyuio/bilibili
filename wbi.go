package bilibili

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"maps"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

const (
	cacheImgKey = "imgKey"
	cacheSubKey = "subKey"
)

var (
	_defaultMixinKeyEncTab = []int{
		46, 47, 18, 2, 53, 8, 23, 32, 15, 50, 10, 31, 58, 3, 45, 35, 27, 43, 5, 49,
		33, 9, 42, 19, 29, 28, 14, 39, 12, 38, 41, 13, 37, 48, 7, 16, 24, 55, 40,
		61, 26, 17, 0, 1, 60, 51, 30, 4, 22, 25, 54, 21, 56, 59, 6, 63, 57, 62, 11,
		36, 20, 34, 44, 52,
	}
)

type Storage interface {
	Set(key string, value any)
	Get(key string) (v any, isSet bool)
}

type MemoryStorage struct {
	data map[string]any
	mu   sync.RWMutex
}

func (impl *MemoryStorage) Set(key string, value any) {
	impl.mu.Lock()
	defer impl.mu.Unlock()

	if impl.data == nil {
		impl.data = make(map[string]any)
	}
	impl.data[key] = value
}

func (impl *MemoryStorage) Get(key string) (v any, isSet bool) {
	impl.mu.RLock()
	defer impl.mu.RUnlock()

	if v, isSet = impl.data[key]; isSet {
		return v, true
	}
	return nil, false
}

// WBI 签名实现
// 如果希望以登录的方式获取则使用 WithCookies or WithRawCookies 设置cookie
// 如果希望以未登录的方式获取 WithCookies(nil) 设置cookie为 nil 即可, 这是 Default 行为
//
//	!!! 使用 WBI 的接口 绝对不可以 set header Referer 会导致失败 !!!
//	!!! 大部分使用 WBI 的接口都需要 set header Cookie !!!
//
// see https://github.com/SocialSisterYi/bilibili-API-collect/blob/master/docs/misc/sign/wbi.md
type WBI struct {
	cookies        []*http.Cookie
	mixinKeyEncTab []int

	// updateCheckerInterval is the interval to check and update wbi keys
	// default is 60 minutes. so if lastInitTime + updateCheckerInterval < now, it will update wbi keys
	updateCheckerInterval time.Duration
	lastInitTime          time.Time
	storage               Storage

	mu      sync.Mutex
	refresh chan struct{}
	http    *resty.Client
	owner   *Client
}

func NewDefaultWbi() *WBI {
	return &WBI{
		cookies:        nil,
		mixinKeyEncTab: _defaultMixinKeyEncTab,

		updateCheckerInterval: 60 * time.Minute,
		storage:               &MemoryStorage{},
		refresh:               make(chan struct{}, 1),
		http:                  resty.New().SetTimeout(20 * time.Second),
	}
}

func (wbi *WBI) WithUpdateInterval(updateInterval time.Duration) *WBI {
	wbi.mu.Lock()
	defer wbi.mu.Unlock()
	wbi.updateCheckerInterval = updateInterval
	return wbi
}

func (wbi *WBI) WithCookies(cookies []*http.Cookie) *WBI {
	wbi.mu.Lock()
	defer wbi.mu.Unlock()
	wbi.cookies = cloneCookies(cookies)
	return wbi
}

func (wbi *WBI) WithRawCookies(rawCookies string) *WBI {
	header := http.Header{}
	header.Add("Cookie", rawCookies)
	req := http.Request{Header: header}

	return wbi.WithCookies(req.Cookies())
}

func (wbi *WBI) WithMixinKeyEncTab(mixinKeyEncTab []int) *WBI {
	wbi.mu.Lock()
	defer wbi.mu.Unlock()
	wbi.mixinKeyEncTab = append([]int(nil), mixinKeyEncTab...)
	return wbi
}

func (wbi *WBI) WithStorage(storage Storage) *WBI {
	wbi.mu.Lock()
	defer wbi.mu.Unlock()
	if storage == nil {
		storage = &MemoryStorage{}
	}
	wbi.storage = storage
	wbi.lastInitTime = time.Time{}
	return wbi
}

func (wbi *WBI) GetKeys(ctx context.Context) (imgKey string, subKey string, err error) {
	return wbi.getKeysContext(ctx)
}

func (wbi *WBI) cachedKeys() (string, string, bool) {
	wbi.mu.Lock()
	defer wbi.mu.Unlock()
	img, sub := wbi.getKeys()
	return img, sub, img != "" && sub != "" && time.Since(wbi.lastInitTime) < wbi.updateCheckerInterval
}

func (wbi *WBI) getKeysContext(ctx context.Context) (string, string, error) {
	if err := checkContext(ctx); err != nil {
		return "", "", err
	}
	if img, sub, valid := wbi.cachedKeys(); valid {
		return img, sub, nil
	}
	// A cancellable gate serializes refreshes without spawning background requests.
	select {
	case wbi.refresh <- struct{}{}:
	case <-ctx.Done():
		return "", "", ctx.Err()
	}
	defer func() { <-wbi.refresh }()
	if err := checkContext(ctx); err != nil {
		return "", "", err
	}
	if img, sub, valid := wbi.cachedKeys(); valid {
		return img, sub, nil
	}
	if err := wbi.doInitWbi(ctx); err != nil {
		return "", "", err
	}
	wbi.mu.Lock()
	defer wbi.mu.Unlock()
	img, sub := wbi.getKeys()
	return img, sub, nil
}

func (wbi *WBI) getKeys() (imgKey string, subKey string) {
	if v, isSet := wbi.storage.Get(cacheImgKey); isSet {
		imgKey, _ = v.(string)
	}

	if v, isSet := wbi.storage.Get(cacheSubKey); isSet {
		subKey, _ = v.(string)
	}

	return imgKey, subKey
}

func (wbi *WBI) SetKeys(imgKey, subKey string) {
	wbi.mu.Lock()
	defer wbi.mu.Unlock()
	wbi.storage.Set(cacheImgKey, imgKey)
	wbi.storage.Set(cacheSubKey, subKey)
	wbi.lastInitTime = time.Now()
}

func (wbi *WBI) GetMixinKey(ctx context.Context) (string, error) { return wbi.mixinKeyContext(ctx) }

func (wbi *WBI) mixinKeyContext(ctx context.Context) (string, error) {
	imgKey, subKey, err := wbi.getKeysContext(ctx)
	if err != nil {
		return "", err
	}

	key := wbi.GenerateMixinKey(imgKey + subKey)
	if len(key) != 32 {
		return "", errors.New("invalid WBI mixin key")
	}
	return key, nil
}

func (wbi *WBI) GenerateMixinKey(orig string) string {
	wbi.mu.Lock()
	defer wbi.mu.Unlock()
	var str strings.Builder
	for _, v := range wbi.mixinKeyEncTab {
		if v >= 0 && v < len(orig) {
			str.WriteByte(orig[v])
		}
	}
	if str.Len() < 32 {
		return ""
	}
	return str.String()[:32]
}

func (wbi *WBI) sanitizeString(s string) string {
	unwantedChars := []string{"!", "'", "(", ")", "*"}
	for _, char := range unwantedChars {
		s = strings.ReplaceAll(s, char, "")
	}
	return s
}

func (wbi *WBI) SignQuery(ctx context.Context, query url.Values, ts time.Time) (url.Values, error) {
	return wbi.signQueryContext(ctx, query, ts)
}

func (wbi *WBI) signQueryContext(ctx context.Context, query url.Values, ts time.Time) (newQuery url.Values, err error) {
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	payload := make(map[string]string, 10)
	for k := range query {
		if len(query[k]) != 1 {
			return nil, errors.New("WBI query requires one value per key")
		}
		payload[k] = query.Get(k)
	}

	newPayload, err := wbi.signMapContext(ctx, payload, ts)
	if err != nil {
		return query, err
	}

	newQuery = url.Values{}
	for k, v := range newPayload {
		newQuery.Set(k, v)
	}

	return newQuery, nil
}

func (wbi *WBI) SignMap(ctx context.Context, payload map[string]string, ts time.Time) (map[string]string, error) {
	return wbi.signMapContext(ctx, payload, ts)
}

func (wbi *WBI) signMapContext(ctx context.Context, payload map[string]string, ts time.Time) (newPayload map[string]string, err error) {
	if err := checkContext(ctx); err != nil {
		return nil, err
	}
	newPayload = maps.Clone(payload)
	if newPayload == nil {
		newPayload = make(map[string]string)
	}
	delete(newPayload, "w_rid")

	newPayload["wts"] = strconv.FormatInt(ts.Unix(), 10)

	// Sort keys
	keys := make([]string, 0, 10)
	for k := range newPayload {
		keys = append(keys, k)
	}

	sort.Strings(keys)

	// Remove unwanted characters
	for k, v := range newPayload {
		v = wbi.sanitizeString(v)
		newPayload[k] = v
	}

	// Build URL parameters
	signQuery := url.Values{}
	for _, k := range keys {
		signQuery.Set(k, newPayload[k])
	}
	signQueryStr := signQuery.Encode()

	// Get mixin key
	mixinKey, err := wbi.mixinKeyContext(ctx)
	if err != nil {
		return payload, err
	}

	// Calculate w_rid
	hash := md5.Sum([]byte(signQueryStr + mixinKey))
	newPayload["w_rid"] = hex.EncodeToString(hash[:])

	return newPayload, nil
}

func (wbi *WBI) doInitWbi(ctx context.Context) error {
	wbi.mu.Lock()
	cookies := cloneCookies(wbi.cookies)
	transport := wbi.http
	owner := wbi.owner
	wbi.mu.Unlock()
	result := struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			WbiImg struct {
				ImgUrl string `json:"img_url"`
				SubUrl string `json:"sub_url"`
			} `json:"wbi_img"`
		}
	}{}

	r := transport.R().SetContext(ctx).SetCookies(cookies)
	if owner != nil {
		r = owner.newRequest(ctx)
	}
	r.
		SetHeader("Accept", "application/json").
		SetHeader("Accept-Language", "zh-CN,zh;q=0.9").
		SetHeader("Origin", "https://www.bilibili.com").
		SetHeader("Referer", "https://www.bilibili.com/").
		SetHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0")
	var resp *resty.Response
	var err error
	if owner != nil {
		resp, err = owner.sendRaw(r, resty.MethodGet, "https://api.bilibili.com/x/web-interface/nav")
	} else {
		resp, err = r.Get("https://api.bilibili.com/x/web-interface/nav")
	}

	if err != nil {
		return errors.WithStack(err)
	}
	if resp.StatusCode() != http.StatusOK {
		return errors.Errorf("status code: %d", resp.StatusCode())
	}

	if err := decodeWBIResponse(resp.Body(), &result); err != nil {
		return err
	}
	if result.Code != 0 {
		if result.Data.WbiImg.ImgUrl == "" || result.Data.WbiImg.SubUrl == "" {
			return errors.Errorf("init wbi 失败, 错误码: %d, 错误信息: %s", result.Code, result.Message)
		}
	}

	if owner == nil && len(resp.Cookies()) > 0 {
		// update cookie
		wbi.WithCookies(resp.Cookies())
	}

	imgKey := strings.Split(strings.Split(result.Data.WbiImg.ImgUrl, "/")[len(strings.Split(result.Data.WbiImg.ImgUrl, "/"))-1], ".")[0]
	subKey := strings.Split(strings.Split(result.Data.WbiImg.SubUrl, "/")[len(strings.Split(result.Data.WbiImg.SubUrl, "/"))-1], ".")[0]

	if len(imgKey) != 32 || len(subKey) != 32 {
		return errors.New("WBI response contains invalid keys")
	}
	wbi.SetKeys(imgKey, subKey)
	return nil
}

// Configuration setters are synchronized; configure custom storage before requests.
func cloneCookies(cookies []*http.Cookie) []*http.Cookie {
	result := make([]*http.Cookie, 0, len(cookies))
	for _, cookie := range cookies {
		if cookie != nil {
			copy := *cookie
			copy.Unparsed = append([]string(nil), cookie.Unparsed...)
			result = append(result, &copy)
		}
	}
	return result
}

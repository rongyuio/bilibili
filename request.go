package bilibili

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/pkg/errors"
)

// Request describes a standard Bilibili code/message/data API request.
// Form and JSON are mutually exclusive. WBI signs query parameters only.
// CSRF parameters must be supplied in the location required by the endpoint.
type Request struct {
	Method  string
	URL     string
	Headers http.Header
	Query   url.Values
	Form    url.Values
	JSON    any
	WBI     bool
}

// Do decodes data into out, which must be a non-nil pointer, or nil to discard data.
// On failure out is unchanged. Client configuration must not change during requests.
func (c *Client) Do(ctx context.Context, req Request, out any) error {
	if ctx == nil {
		return errors.New("request context is nil")
	}
	if err := validateOutput(out); err != nil {
		return err
	}
	u, err := url.Parse(req.URL)
	if err != nil || u.Host == "" || (u.Scheme != "https" && u.Scheme != "http") || u.User != nil {
		return errors.New("request URL must be an absolute HTTP(S) URL without user information")
	}
	if req.Method == "" {
		return errors.New("request method is empty")
	}
	if req.Form != nil && req.JSON != nil {
		return errors.New("Form and JSON cannot both be set")
	}
	r := c.newRequest(ctx)
	r.Header = req.Headers.Clone()
	if r.Header == nil {
		r.Header = make(http.Header)
	}
	// Move URL query parameters into the request before signing, including duplicates.
	r.QueryParam, err = url.ParseQuery(u.RawQuery)
	if err != nil {
		return errors.New("request URL contains an invalid query")
	}
	for key, values := range req.Query {
		r.QueryParam[key] = append([]string(nil), values...)
	}
	u.RawQuery = ""
	u.Fragment = ""
	if req.Form != nil {
		r.SetFormDataFromValues(req.Form)
		r.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	}
	if req.JSON != nil {
		body, err := json.Marshal(req.JSON)
		if err != nil {
			return fmt.Errorf("encode JSON request body: %w", err)
		}
		r.SetBody(body)
		r.SetHeader("Content-Type", "application/json")
	}
	if req.WBI {
		if err := c.fillWbi()(r); err != nil {
			return err
		}
	}
	return c.send(r, req.Method, u.String(), out)
}

func validateOutput(out any) error {
	if out != nil {
		v := reflect.ValueOf(out)
		if v.Kind() != reflect.Pointer || v.IsNil() {
			return errors.New("output must be a non-nil pointer or nil")
		}
	}
	return nil
}

func (c *Client) fillWbi() paramHandler {
	return func(r *resty.Request) error {
		// Resty merges client defaults later; include them before calculating w_rid.
		for key, values := range c.resty.QueryParam {
			if _, exists := r.QueryParam[key]; !exists {
				r.QueryParam[key] = append([]string(nil), values...)
			}
		}
		query, err := c.wbi.signQueryContext(r.Context(), r.QueryParam, time.Now())
		if err != nil {
			return fmt.Errorf("sign WBI query: %w", err)
		}
		r.QueryParam = query
		// Cookies were captured when the request was created.
		// An empty value overrides a Referer inherited from the Resty client.
		r.SetHeader("Referer", "")
		return nil
	}
}

func execute[Out any](c *Client, method, endpoint string, in any, handlers ...paramHandler) (out Out, err error) {
	return executeRequest[Out](c, c.newRequest(context.Background()), method, endpoint, in, handlers...)
}

func executeRequest[Out any](c *Client, r *resty.Request, method, endpoint string, in any, handlers ...paramHandler) (out Out, err error) {
	if err = withParams(r, in); err != nil {
		return out, err
	}
	for _, handler := range handlers {
		if err = handler(r); err != nil {
			return out, err
		}
	}
	err = c.send(r, method, endpoint, &out)
	return out, err
}

func (c *Client) send(r *resty.Request, method, endpoint string, out any) error {
	resp, err := c.sendRaw(r, method, endpoint)
	if err != nil {
		return err
	}
	if resp.StatusCode() != http.StatusOK {
		return fmt.Errorf("%s %s: HTTP status %d", method, safeEndpoint(endpoint), resp.StatusCode())
	}
	return decodeResponse(method, endpoint, resp.Body(), out)
}

type paramHandler func(*resty.Request) error

func fillCsrf(_ *Client) paramHandler {
	return func(r *resty.Request) error {
		csrf := cookieValue(r.Cookies, "bili_jct")
		if len(csrf) == 0 {
			return errors.New("B站登录过期")
		}
		r.SetQueryParam("csrf", csrf)
		r.SetQueryParam("csrf_token", csrf)
		return nil
	}
}

func fillParam(key, value string) paramHandler {
	return func(r *resty.Request) error {
		r.SetQueryParam(key, value)
		return nil
	}
}

// newRequest snapshots cookies once. Configuration is immutable while requests run.
func (c *Client) newRequest(ctx context.Context) *resty.Request {
	return c.resty.R().SetContext(ctx).SetCookies(c.GetCookies())
}

// sendRaw merges response cookies even when the status or business code is an error.
func (c *Client) sendRaw(r *resty.Request, method, endpoint string) (*resty.Response, error) {
	resp, err := r.Execute(method, endpoint)
	if resp != nil && resp.RawResponse != nil {
		c.SetCookies(resp.Cookies())
	}
	if err != nil {
		return resp, fmt.Errorf("%s %s: %w", method, safeEndpoint(endpoint), err)
	}
	if resp == nil || resp.RawResponse == nil {
		return nil, fmt.Errorf("%s %s: missing HTTP response", method, safeEndpoint(endpoint))
	}
	return resp, nil
}

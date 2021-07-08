package binanceapi

import (
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

type Headers struct {
	ClientType      string
	Cookie          string
	CsrfToken       string
	UserAgent       string
}

func (a *Api) postRequest(url string, b []byte) (*fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	a.headers.initHeaders(req)
	req.Header.SetMethod("POST")
	req.Header.SetRequestURI(url)
	req.Header.SetContentType("application/json")
	req.SetBody(b)
	resp := fasthttp.AcquireResponse()
	client := fasthttp.Client{}
	if a.proxy != "" {
		client.Dial = fasthttpproxy.FasthttpHTTPDialer(a.proxy)
	}
	if err := client.Do(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (h *Headers) initHeaders(r *fasthttp.Request) {
	r.Header.Set("clienttype", h.ClientType)
	r.Header.Set("cookie", h.Cookie)
	r.Header.Set("csrftoken", h.CsrfToken)
	r.Header.Set("user-agent", h.UserAgent)
}
package binanceapi

import (
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

type Headers struct {
	ClientType string
	Cookie     string
	CsrfToken  string
	UserAgent  string
}

func (a *Api) GenerateRequest(url string, b []byte) *fasthttp.Request {
	req := fasthttp.AcquireRequest()
	a.headers.initHeaders(req)
	req.Header.SetMethod("POST")
	req.Header.SetRequestURI(url)
	req.Header.SetContentType("application/json")
	req.SetBody(b)
	return req
}

func (a *Api) get(url string) (*fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	req.Header.SetMethod("GET")
	req.Header.SetRequestURI(url)
	resp := fasthttp.AcquireResponse()
	if err := fasthttp.Do(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *Api) GenerateHttpClient() {
	client := &fasthttp.Client{}
	if a.proxy != "" {
		client.Dial = fasthttpproxy.FasthttpHTTPDialer(a.proxy)
	}
	a.httpclient.client = client
}

func (a *Api) postRequest(req *fasthttp.Request) (*fasthttp.Response, error) {
	resp := fasthttp.AcquireResponse()
	if err := a.httpclient.client.Do(req, resp); err != nil {
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

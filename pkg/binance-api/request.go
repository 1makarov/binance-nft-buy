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

const buyurl = "https://www.binance.com/bapi/nft/v1/private/nft/mystery-box/purchase"

func (a *Api) GenerateRequest(b []byte) (*fasthttp.Request, *fasthttp.Client) {
	req := fasthttp.AcquireRequest()
	a.headers.initHeaders(req)
	req.Header.SetMethod("POST")
	req.Header.SetRequestURI(buyurl)
	req.Header.SetContentType("application/json")
	client := fasthttp.Client{}
	if a.proxy != "" {
		client.Dial = fasthttpproxy.FasthttpHTTPDialer(a.proxy)
	}
	req.SetBody(b)
	return req, &client
}

func (a *Api) postRequest(req *fasthttp.Request, client *fasthttp.Client) (*fasthttp.Response, error) {
	resp := fasthttp.AcquireResponse()
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

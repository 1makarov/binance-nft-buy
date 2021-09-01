package bapi

import (
	"github.com/valyala/fasthttp"
)


func (api *Api) get(url string) (*fasthttp.Response, error) {
	r := fasthttp.AcquireRequest()
	api.request.CopyTo(r)
	r.Header.SetRequestURI(url)
	r.Header.SetMethod(fasthttp.MethodGet)
	resp := fasthttp.AcquireResponse()
	if err := api.http.Do(r, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func getPublic(url string) (*fasthttp.Response, error) {
	r := fasthttp.AcquireRequest()
	r.Header.SetRequestURI(url)
	r.Header.SetMethod(fasthttp.MethodGet)
	resp := fasthttp.AcquireResponse()
	if err := fasthttp.Do(r, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *Api) post(url string, body []byte) (*fasthttp.Response, error) {
	req := fasthttp.AcquireRequest()
	api.request.CopyTo(req)
	req.Header.SetRequestURI(url)
	req.Header.SetMethod(fasthttp.MethodPost)
	req.Header.SetContentType("application/json")
	req.SetBody(body)
	resp := fasthttp.AcquireResponse()
	if err := api.http.Do(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

func (api *Api) postRequest(req *fasthttp.Request) (*fasthttp.Response, error) {
	resp := fasthttp.AcquireResponse()
	if err := api.http.Do(req, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

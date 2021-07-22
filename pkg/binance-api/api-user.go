package binanceapi

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

func (a *Api) UserId() (*userresponse, error) {
	req := a.GenerateRequest(URLinfo, nil)
	client := a.GenerateHttpClient()
	resp, err := a.postRequest(req, client)
	if err = handleError(resp, err); err != nil {
		return nil, err
	}
	user, err := a.unmarshalUser(resp)
	if err != nil {
		return nil, err
	}
	return user, nil
}

type userresponse struct {
	Data struct {
		Agentid uint64 `json:"agentId"`
	} `json:"data"`
}

func (a *Api) unmarshalUser(resp *fasthttp.Response) (*userresponse, error) {
	var response userresponse
	err := json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

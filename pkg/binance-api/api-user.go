package binanceapi

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

func (a *Api) User() (*userResponse, error) {
	req := a.GenerateRequest(URLInfo, nil)
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

type userResponse struct {
	Data struct {
		Email string `json:"email"`
	} `json:"data"`
}

func (a *Api) unmarshalUser(resp *fasthttp.Response) (*userResponse, error) {
	var response userResponse
	err := json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

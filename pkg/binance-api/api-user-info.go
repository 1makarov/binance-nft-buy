package bapi

import (
	"encoding/json"
	binance_struct "github.com/1makarov/binance-nft-buy/internal/domain/binance-api"
	"github.com/valyala/fasthttp"
)

const (
	urlUserInfo = "https://www.binance.com/bapi/accounts/v1/private/account/user/base-detail"
)

func (api *Api) GetEmail() (string, error) {
	resp, err := api.post(urlUserInfo, nil)
	if err = handleError(resp, err); err != nil {
		return "", err
	}
	user, err := unmarshalUserInfo(resp)
	if err != nil {
		return "", err
	}
	return user.Data.Email, nil
}

func unmarshalUserInfo(resp *fasthttp.Response) (*binance_struct.UserInformationResponse, error) {
	var response binance_struct.UserInformationResponse
	err := json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

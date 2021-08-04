package binanceapi

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
)

func (a *Api) NFTInfo() (*NftInfoResponse, error) {
	resp, err := a.get(URLNftInfo)
	if err = handleError(resp, err); err != nil {
		return nil, err
	}
	info, err := a.unmarshalNFTInfo(resp)
	if err != nil {
		return nil, err
	}
	return info, nil
}

type NftInfoResponse struct {
	Boxes []struct {
		Box
	} `json:"data"`
}

type Box struct {
	Name          string `json:"name"`
	Productid     string `json:"productId"`
	StartTime     int64  `json:"startTime"`
	MappingStatus int    `json:"mappingStatus"`
}

func (a *Api) unmarshalNFTInfo(resp *fasthttp.Response) (*NftInfoResponse, error) {
	var response NftInfoResponse
	err := json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

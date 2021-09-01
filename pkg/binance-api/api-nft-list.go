package bapi

import (
	"encoding/json"
	binance_struct "github.com/1makarov/binance-nft-buy/internal/domain/binance-api"
	"github.com/valyala/fasthttp"
)

const (
	urlNFTMysteryBoxList = "https://www.binance.com/bapi/nft/v1/public/nft/mystery-box/list?page=1&size=15"
)

func NFTMysteryBoxList() (*binance_struct.NftMysteryBoxesListResponse, error) {
	r, err := getPublic(urlNFTMysteryBoxList)
	if err = handleError(r, err); err != nil {
		return nil, err
	}
	b, err := unmarshalNFTMysteryBoxList(r)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalNFTMysteryBoxList(resp *fasthttp.Response) (*binance_struct.NftMysteryBoxesListResponse, error) {
	var response binance_struct.NftMysteryBoxesListResponse
	err := json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

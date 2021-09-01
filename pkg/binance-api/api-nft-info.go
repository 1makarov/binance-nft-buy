package bapi

import (
	"encoding/json"
	"fmt"
	binance_struct "github.com/1makarov/binance-nft-buy/internal/domain/binance-api"
	"github.com/1makarov/binance-nft-buy/internal/domain/mysterybox"
	"github.com/valyala/fasthttp"
	"strconv"
)


const (
	urlNFTMysteryBoxInfo = "https://www.binance.com/bapi/nft/v1/friendly/nft/mystery-box/detail?productId=%s"
)

func NFTMysteryBoxInfo(productID string) (*mysterybox.Information, error) {
	r, err := getPublic(fmt.Sprintf(urlNFTMysteryBoxInfo, productID))
	if err = handleError(r, err); err != nil {
		return nil, err
	}
	b, err := unmarshalNFTMysteryBoxInfo(r)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func unmarshalNFTMysteryBoxInfo(resp *fasthttp.Response) (*mysterybox.Information, error) {
	var response binance_struct.NftMysteryBoxesInfoResponse
	err := json.Unmarshal(resp.Body(), &response)
	if err != nil {
		return nil, err
	}
	price, err := strconv.ParseFloat(response.Data.Price, 64)
	if err != nil {
		return nil, err
	}
	userBalance, err := strconv.ParseFloat(response.Data.UserBalance, 64)
	if err != nil {
		return nil, err
	}
	return &mysterybox.Information{
		Price:       price,
		Balance:     userBalance,
		LimitPerBuy: response.Data.LimitPerTime,
		StartTime:   response.Data.StartTime / 1000,
	}, nil
}

package binance

import (
	binanceapi "github.com/1makarov/binance-nft-buy/pkg/binance-api"
)

type Client struct {
	api *binanceapi.Api
}

func CreateClient(api *binanceapi.Api) *Client {
	return &Client{
		api: api,
	}
}

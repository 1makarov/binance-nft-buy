package binance

import (
	"fmt"
	binanceapi "github.com/1makarov/binance-nft-buy/pkg/binance-api"
)

var (
	errorNoSale = fmt.Errorf("There are no sales at the moment")
)

type Client struct {
	api *binanceapi.Api
}

func CreateClient(api *binanceapi.Api) *Client {
	return &Client{
		api: api,
	}
}

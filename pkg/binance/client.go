package binance

import (
	binanceapi "github.com/1makarov/binance-nft-buy/pkg/binance-api"
)

type Client struct {
	api        *binanceapi.Api
	mysteryBox *binanceapi.MysteryBox
	time       string
}

func CreateClient(api *binanceapi.Api, mysteryBox *binanceapi.MysteryBox, time string) *Client {
	return &Client{
		api:        api,
		mysteryBox: mysteryBox,
		time:       time,
	}
}

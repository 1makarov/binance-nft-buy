package binance

import (
	binanceapi "github.com/1makarov/binance-nft-buy/pkg/binance-api"
	"log"
	"time"
)

type Client struct {
	api        *binanceapi.Api
	mysteryBox []byte
	time       time.Time
}

func CreateClient(api *binanceapi.Api, mysteryBox MysteryBox, time string) *Client {
	return &Client{
		api:        api,
		mysteryBox: jsonMysteryBox(&mysteryBox),
		time:       initTime(time),
	}
}

const layout = "02/01/2006 15:04:05"

func initTime(i string) time.Time {
	t, err := time.Parse(layout, i)
	if err != nil {
		log.Fatalln(err)
	}
	return t.UTC()
}

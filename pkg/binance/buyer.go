package binance

import (
	"fmt"
	binanceapi "github.com/1makarov/binance-nft-buy/pkg/binance-api"
	"log"
	"time"
)

func (c *Client) Start() {
	defer fmt.Scanf("\n")
	// generate request to get info
	req := c.api.GenerateRequest(binanceapi.URLinfo, nil)
	// get user id (check cookie)
	user, err := c.api.UserId()
	if err != nil {
		log.Fatalln(err)
	}
	// notification for cookie
	notification(user.Data.Agentid)
	// generate httpclient
	httpclient := c.api.GenerateHttpClient()
	// generate bytes from json
	box := binanceapi.MarshalBoxBuy(c.mysteryBox)
	// generate request to buy box
	req = c.api.GenerateRequest(binanceapi.URLbuy, *box)
	// handle time from config
	t := c.handleTime()
	// wait buy time
	waitBuyTime(t)

	if err = c.api.MysteryBoxBuy(req, httpclient); err != nil {
		log.Fatalln(err)
	}
}

func notification(id uint64) {
	log.Println(fmt.Sprintf("You have entered working cookies, your ID: %d", id))
}

func waitBuyTime(t time.Time) {
	for {
		if time.Now().UTC().After(t) {
			return
		}
	}
}

const layout = "02/01/2006 15:04:05"

func (c *Client) handleTime() time.Time {
	t, err := time.Parse(layout, c.time)
	if err != nil {
		log.Fatalln(err)
	}
	return t.UTC()
}

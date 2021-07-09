package binance

import (
	"log"
	"time"
)

func (c *Client) Start() {
	req, httpclient := c.api.GenerateRequest(c.mysteryBox)

	c.waitSell()

	if err := c.api.MysteryBoxBuy(req, httpclient); err != nil {
		log.Fatalln(err)
	}
}

func (c *Client) waitSell() {
	for {
		if time.Now().UTC().After(c.time) {
			return
		}
	}
}

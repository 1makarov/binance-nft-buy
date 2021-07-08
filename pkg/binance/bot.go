package binance

import (
	"log"
	"time"
)

func (c *Client) Start() {
	c.waitSell()

	if err := c.api.MysteryBoxBuy(c.mysteryBox); err != nil {
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
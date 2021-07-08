package binance

import (
	"encoding/json"
	"log"
)

type MysteryBox struct {
	Volume    int    `json:"number"`
	Productid string `json:"productId"`
}

func jsonMysteryBox(box *MysteryBox) []byte {
	b, err := json.Marshal(box)
	if err != nil {
		log.Fatalln(err)
	}
	return b
}

package binanceapi

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

func (a *Api) MysteryBoxBuy(req *fasthttp.Request, client *fasthttp.Client) error {
	resp, err := a.postRequest(req, client)
	if err = handleError(resp, err); err != nil {
		return err
	}
	log.Println(fmt.Sprintf("Information about buying: %s", string(resp.Body())))
	return nil
}

type MysteryBox struct {
	Volume    int    `json:"number"`
	ProductId string `json:"productId"`
}

func MarshalBoxBuy(box *MysteryBox) *[]byte {
	b, err := json.Marshal(box)
	if err != nil {
		log.Fatalln(err)
	}
	return &b
}
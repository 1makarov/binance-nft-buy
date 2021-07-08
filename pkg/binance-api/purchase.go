package binanceapi

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

const (
	mysteryBoxBuyURL      = "https://www.binance.com/bapi/nft/v1/private/nft/mystery-box/purchase"
	errorStatusCodeNot200 = "status code %d, %s"
)

func (a *Api) MysteryBoxBuy(body []byte) error {
	resp, err := a.postRequest(mysteryBoxBuyURL, body)
	if err != nil {
		return err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		return fmt.Errorf(errorStatusCodeNot200, resp.StatusCode(), string(resp.Body()))
	}
	fmt.Println(resp)
	return nil
}

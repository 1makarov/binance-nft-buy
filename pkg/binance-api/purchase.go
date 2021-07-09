package binanceapi

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

const errorStatusCodeNot200 = "status code %d, %s"

func (a *Api) MysteryBoxBuy(req *fasthttp.Request, client *fasthttp.Client) error {
	resp, err := a.postRequest(req, client)
	if err != nil {
		return err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		return fmt.Errorf(errorStatusCodeNot200, resp.StatusCode(), string(resp.Body()))
	}
	fmt.Println(resp)
	return nil
}

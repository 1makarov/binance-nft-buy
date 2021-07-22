package binanceapi

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func handleError(resp *fasthttp.Response, err error) error {
	if err != nil {
		return err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		return fmt.Errorf(errorStatusCodeNot200, resp.StatusCode(), string(resp.Body()))
	}
	return nil
}

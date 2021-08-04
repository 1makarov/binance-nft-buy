package binanceapi

import "github.com/valyala/fasthttp"

const (
	URLBuy     = "https://www.binance.com/bapi/nft/v1/private/nft/mystery-box/purchase"
	URLInfo    = "https://www.binance.com/bapi/accounts/v1/private/account/user/base-detail"
	URLNftInfo = "https://www.binance.com/bapi/nft/v1/public/nft/mystery-box/list?page=1&size=15"
)

type Api struct {
	proxy   string
	headers *Headers
	httpclient
}

type httpclient struct {
	client *fasthttp.Client
}

func CreateApi(proxy string, headers *Headers) *Api {
	return &Api{
		proxy:   proxy,
		headers: headers,
	}
}

package binanceapi

const (
	URLbuy  = "https://www.binance.com/bapi/nft/v1/private/nft/mystery-box/purchase"
	URLinfo = "https://www.binance.com/bapi/accounts/v1/private/account/user/base-detail"

	errorStatusCodeNot200 = "status code %d, %s"
)

type Api struct {
	proxy   string
	headers *Headers
}

func CreateApi(proxy string, headers *Headers) *Api {
	return &Api{
		proxy:   proxy,
		headers: headers,
	}
}

package binanceapi

type Api struct {
	proxy string
	headers Headers
}

func CreateApi(proxy string, headers Headers) *Api {
	return &Api{
		proxy: proxy,
		headers: headers,
	}
}

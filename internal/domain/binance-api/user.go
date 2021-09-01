package binance_api

type UserInformationResponse struct {
	Data UserInfoResponse `json:"data"`
}

type UserInfoResponse struct {
	Email string `json:"email"`
}

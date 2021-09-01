package mysterybox

type Information struct {
	Price       float64
	Balance     float64
	LimitPerBuy uint64
	StartTime   int64
}

type Box struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

package structure

type Token struct {
	Symbol       string  `json:"symbol"`
	Image        string  `json:"image"`
	CurrentPrice float64 `json:"current_price"`
	MarketCap    float64 `json:"market_cap"`
}

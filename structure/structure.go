package structure

type Token struct {
	Symbol                        string  `json:"symbol"`
	Image                         string  `json:"image"`
	CurrentPrice                  float64 `json:"current_price"`
	MarketCap                     float64 `json:"market_cap"`
	Price_change_percentage_24h   float64 `json:"price_change_percentage_24h"`
	FormattedPrice_percentage_24h string
	Id                            int
	FormattedMarketCap            string
	IsPricePercentagePositive     bool
}

type DataReceived struct {
	Address string `json:"address"`
}

type UserData struct {
	LiveUser string `json:"liveuser"`
	Address  string `json:"address"`
}

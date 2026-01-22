package structure

type Token struct { // Représente un jeton dans la liste globale avec les tags JSON pour CoinGecko et des champs formatés pour le front.
	FullName                      string  `json:"id"`
	Symbol                        string  `json:"symbol"`
	Image                         string  `json:"image"`
	CurrentPrice                  float64 `json:"current_price"`
	MarketCap                     float64 `json:"market_cap"`
	Price_change_percentage_24h   float64 `json:"price_change_percentage_24h"`
	FormattedPrice_percentage_24h string
	Id                            int
	FormattedMarketCap            string
	IsPricePercentagePositive     bool
	Type                          string
	IsFavorite                    bool
}

type TokenInfo struct { // Structure complexe pour les détails d'un jeton, incluant les données imbriquées de l'API (MarketData, Tickers).
	FullName    string `json:"name"`
	Description struct {
		En string `json:"en"`
	} `json:"description"`
	Links struct {
		Homepage []string `json:"homepage"`
	} `json:"links"`
	Tickers []struct {
		ConvertedVolume struct {
			USD float64 `json:"usd"`
		} `json:"converted_volume"`
	} `json:"tickers"`
	MarketData       MarketData `json:"market_data"`
	Supply           string
	MarketCap        string
	DescriptionFinal string
	link             string
	VolumeUSD        string
	ImgUrl           struct {
		Large string `json:"large"`
	} `json:"image"`
	Image  string
	WebUrl string
}

type DataReceived struct { // Structure temporaire pour le décodage du corps de la requête lors de la réception d'une adresse wallet.
	Address string `json:"address"`
}

type UserData struct { // Format de stockage persistant pour l'utilisateur actuel et son adresse associée.
	LiveUser string `json:"liveuser"`
	Address  string `json:"address"`
}

type MarketData struct { // Contient les métriques d'approvisionnement (Supply) et de capitalisation boursière brute.
	TotalSupply       float64 `json:"total_supply"`
	MaxSupply         float64 `json:"max_supply"`
	MaxSupplyInfinite bool    `json:"max_supply_infinite"`
	CirculatingSupply float64 `json:"circulating_supply"`
	LastUpdated       string  `json:"last_updated"`
	MarketCap         struct {
		USD float64 `json:"usd"`
	} `json:"market_cap"`
}

type Filters struct { // Définit l'état des filtres sélectionnés via les cases à cocher du formulaire.
	Layer1   bool
	Layer2   bool
	Memecoin bool
}

type Data struct { // Structure de données racine envoyée au template pour le rendu des listes et l'état des filtres.
	Tokens  []Token
	Filters Filters
}

package types

type Ticker struct {
	CurrentPrice float64 `json:"current_price"`
	TickerName   string  `json:"ticker_name"`
}

type ApiResponse struct {
	Tickers []Ticker `json:"tickers"`
}

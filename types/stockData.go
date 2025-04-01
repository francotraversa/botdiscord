package types

type Indicators struct {
	SMA50      float64 `json:"sma50"`
	SMA200     float64 `json:"sma200"`
	RSI        float64 `json:"rsi"`
	EMA50      float64 `json:"ema50"`
	EMA200     float64 `json:"ema200"`
	BB_BBM     float64 `json:"bb_bbm"`
	BB_BBH     float64 `json:"bb_bbh"`
	BB_BBL     float64 `json:"bb_bbl"`
	MACD       float64 `json:"macd"`
	MACDSignal float64 `json:"macd_signal"`
	MACDDiff   float64 `json:"macd_diff"`
}

type StockData struct {
	Ticker     string     `json:"ticker"`
	Close      float64    `json:"close"`
	Decision   string     `json:"decision"`
	Score      float64    `json:"score"`
	Indicators Indicators `json:"indicators"`
}

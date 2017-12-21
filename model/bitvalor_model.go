package model


type BitValorResponse struct {
	Ticker24h Ticker `json:"ticker_24h"`
	Ticker12h Ticker `json:"ticker_12h"`
	Ticker1h Ticker `json:"ticker_1h"`	
}

type Ticker struct {
	Total     Values         `json: "total"`
	Exchanges ExchangeValues `json:"exchanges"`
}

type ExchangeValues struct {
	FOX Values
	ARN Values
	B2U Values
	BTD Values
	FLW Values
	LOC Values
	MBT Values
	NEG Values
}

type Values struct {
	Last float64 `json:"last"`
	Open float64 `json:"open"`
	High float64 `json:"high"`
	Low float64 `json:"low"`
	Vol float64 `json:"vol"`
	Vwap float64 `json:"vwap"`
	Money float64 `json:"money"`
	Trades uint64 `json:"trades"`
}
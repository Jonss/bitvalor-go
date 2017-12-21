package model

type Exchange struct {
	Values ExchangeInfo `json:"exchange"`
}


type ExchangeInfo struct {
	Name string `json:"name"`
	Last int64 `json:"last"`
	Open int64 `json:"open"`
	High int64 `json:"high"`
	Low int64 `json:"low"`
	Vol int64 `json:"vol"`
	Vwap int64 `json:"vwap"`
	Money int64 `json:"money"`
	Trades int64 `json:"trades"`
}



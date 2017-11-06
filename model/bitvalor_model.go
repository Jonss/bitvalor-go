package model

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

const BITVALOR_API_URL = "https://api.bitvalor.com/v1/ticker.json"

type BitValorResponse struct {
	Ticker24h Ticker `json:"ticker_24h"`
	Ticker12h Ticker `json:"ticker_12h"`
	Ticker1h Ticker `json:"ticker_1h"`	
}

type Ticker struct {
	Total Values `json: "total"`
	Exchanges Exchanges `json:"exchanges"`
}

type Exchanges struct {
	FOX Values `json:"FOX"`
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

func FetchData() BitValorResponse {
	resp, err := http.Get(BITVALOR_API_URL)
	if err != nil {
		fmt.Println(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	bitValorResponse := BitValorResponse{}
	json.Unmarshal([]byte(string(b)), &bitValorResponse)

	return bitValorResponse
}

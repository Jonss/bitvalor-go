package mapper

import "github.com/jonss/bitvalor-go/model"

func mapBitvalorExchangeToExchange(v model.Values, name string) model.Exchange {
	exchange := model.Exchange{}
	e := model.ExchangeInfo{}
	e.Last = int64(v.Last * 100.0)
	e.Open = int64(v.Open * 100.0)
	e.High = int64(v.High * 100.0)
	e.Low = int64(v.Low * 100.0)
	e.Vol = int64(v.Vol * 100.0)
	e.Vwap = int64(v.Vwap * 100.0)
	e.Money = int64(v.Money * 100.0)
	e.Trades = int64(v.Trades * 100.0)
	e.Name = name
	exchange.Values = e
	return exchange
}


func MapExchanges(exchanges model.ExchangeValues) []model.Exchange {
	i := []model.Exchange{}
	i = append(i, mapBitvalorExchangeToExchange(exchanges.ARN, "arena_bitcoin"))
	i = append(i, mapBitvalorExchangeToExchange(exchanges.FOX, "foxbit"))
	i = append(i, mapBitvalorExchangeToExchange(exchanges.B2U, "bitcoin_to_you"))
	i = append(i, mapBitvalorExchangeToExchange(exchanges.BTD, "bitcoin_trade"))
	i = append(i, mapBitvalorExchangeToExchange(exchanges.FLW, "flow_btc"))
	i = append(i, mapBitvalorExchangeToExchange(exchanges.LOC, "local_bitcoins"))
	i = append(i, mapBitvalorExchangeToExchange(exchanges.MBT,"mercado_bitcoin"))
	i = append(i, mapBitvalorExchangeToExchange(exchanges.NEG, "negocie_bitcoins"))
	return i
}

package main

import (
	"fmt"
	"github.com/jonss/bitvalor-go/request"
	_ "github.com/jonss/bitvalor-go/model"
	"encoding/json"
	"github.com/jonss/bitvalor-go/mapper"
)

func main() {
	
	bitValorResponse := bitvalor_request.FetchData()

	exchanges := bitValorResponse.Ticker1h.Exchanges

	mapExchanges := mapper.MapExchanges(exchanges)

	bytes, _ := json.Marshal(mapExchanges)

	fmt.Println(string(bytes))
}


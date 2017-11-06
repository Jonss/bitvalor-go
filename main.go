
package main

import (
	"fmt"
	"github.com/jonss/bitvalor-go/model"
)

func main() {
	
	bitValorResponse := model.FetchData()

	exchanges := bitValorResponse.Ticker1h.Exchanges

	fmt.Println("Foxbit: ",  exchanges.FOX)
}


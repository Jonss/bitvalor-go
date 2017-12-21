package bitvalor_request

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"github.com/jonss/bitvalor-go/model"
	"encoding/json"
)

const BITVALOR_API_URL = "https://api.bitvalor.com/v1/ticker.json"


func FetchData() model.BitValorResponse {
	resp, err := http.Get(BITVALOR_API_URL)
	if err != nil {
		fmt.Println(err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	bitValorResponse := model.BitValorResponse{}
	json.Unmarshal([]byte(string(b)), &bitValorResponse)

	return bitValorResponse
}

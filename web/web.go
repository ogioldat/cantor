package web

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ogioldat/cantor/config"
	"github.com/ogioldat/cantor/types"
)

var cfg, helpers = config.GetConfig()

func GetLatestCurrencies() {
	for index := range cfg.Endpoints.Currencies {
		param := cfg.Endpoints.Currencies[index]
		var response types.CurrenciesResponse

		requestData(helpers.ApplyURL(param), response)
	}
}

func requestData(query_url string, response interface{}) {
	resp, _ := http.Get(query_url)
	body, _ := ioutil.ReadAll(resp.Body)

	parse_err := json.Unmarshal(body, &response)

	if parse_err != nil {
		panic(parse_err)
	}

	fmt.Println(response)
}

package web

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/ogioldat/cantor/config"
	"github.com/ogioldat/cantor/types"
	"github.com/ogioldat/cantor/utils"
)

var cfg, helpers = config.GetConfig()

func GetCurrencies(date string) types.LatestCurrencies {
	var results []types.CurrenciesResponse
	parse := utils.ParseLatestCurrencies

	for index := range cfg.Endpoints.Currencies {
		var response []types.CurrenciesResponse
		param := cfg.Endpoints.Currencies[index]
		data := requestData(helpers.ApplyURL(param + "/" + date))

		parse_err := json.Unmarshal(data, &response)

		if parse_err != nil {
			continue
		}

		results = append(results, response[0])
	}
	parsed_results := parse(results)

	if len(parsed_results.Items) == 0 {
		return types.LatestCurrencies{
			EffectiveDate: date,
			Items:         []types.LatestCurrenciesItem{},
		}
	}

	return parsed_results
}

func requestData(query_url string) []byte {
	resp, _ := http.Get(query_url)
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

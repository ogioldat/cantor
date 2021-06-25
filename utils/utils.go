package utils

import (
	"net/url"

	"github.com/ogioldat/cantor/types"
)

func ParseURLParams(params map[string]string) string {
	query := url.Values{}

	for param, val := range params {
		query.Add(param, val)
	}

	return query.Encode()
}

func ParseLatestCurrencies(data []types.CurrenciesResponse) types.LatestCurrencies {
	var items []types.LatestCurrenciesItem
	var date string

	for _, datum := range data {
		date = datum.EffectiveDate

		for _, item := range datum.Rates {
			items = append(items, types.LatestCurrenciesItem{
				Name:  item.Currency,
				Code:  item.Code,
				Value: item.Mid,
			})
		}
	}

	return types.LatestCurrencies{
		EffectiveDate: date,
		Items:         items,
	}
}

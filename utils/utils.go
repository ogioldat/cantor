package utils

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

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

func GetDate() string {
	var year, month, day int

	currentYear, currentMonth, currentDay := time.Now().Date()

	for {
		fmt.Println("Podaj rok: (np: 2020)")
		fmt.Scanln(&year)

		if !Between(year, 1900, currentYear) {
			fmt.Println("Niepoprawny rok!")
		} else {
			break
		}
	}

	for {
		fmt.Println("Podaj miesiac: (np: 04)")
		fmt.Scanln(&month)

		if !Between(month, 01, 12) || (year == int(currentYear) && month > int(currentMonth)) {
			fmt.Println("Niepoprawny miesiac!")
		} else {
			break
		}
	}

	for {
		fmt.Println("Podaj dzień: (np: 01)")
		fmt.Scanln(&day)

		if !Between(day, 01, 31) ||
			(year == int(currentYear) && month > int(currentMonth) && day > int(currentDay)) {
			fmt.Println("Niepoprawny dzień!")
		} else {
			break
		}
	}

	return strconv.Itoa(year) + "-" + FillZero(strconv.Itoa(month)) + "-" + FillZero(strconv.Itoa(day))
}

func Between(val int, a int, b int) bool {
	return val >= a && val <= b
}

func FillZero(val string) string {
	parsed, _ := strconv.Atoi(val)
	if parsed < 10 {
		return "0" + val
	}
	return val
}

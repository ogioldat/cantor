package functionalities

import (
	"strconv"

	"github.com/ogioldat/cantor/ui/table"
	"github.com/ogioldat/cantor/web"
)

func DisplayCurrencies(date string) {
	currencies := web.GetCurrencies(date)

	var rows [][]string

	for _, currency := range currencies.Items {
		val := strconv.FormatFloat(float64(currency.Value), 'f', 4, 64)

		rows = append(rows, []string{
			currency.Name,
			val,
			currency.Code,
		})
	}

	table.ShowMyPrettyTable(rows)
}

func DisplayCurrentCurrenciesBuySell() {

}

func DisplayCurrentCurrenciesBuySellOnDay() {

}

func CurrencyConversion() {

}

func Settings() {

}

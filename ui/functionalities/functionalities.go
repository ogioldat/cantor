package functionalities

import (
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/ogioldat/cantor/ui/table"
	"github.com/ogioldat/cantor/web"
)

func DisplayLatestCurrencies() {
	currencies := web.GetLatestCurrencies()
func DisplayCurrentCurrencies() {

	var rows [][]string

	for _, currency := range currencies.Items {
		fmt.Println(currency)
		val := strconv.FormatFloat(float64(currency.Value), 'f', 4, 64)

		rows = append(rows, []string{
			currency.Name,
			val,
			currency.Code,
		})
	}

	table.ShowMyPrettyTable(rows)
}

func DisplayCurrenciesOnDay() {
	fmt("Podaj rok: (np:2020)")
	var year string
    fmt.Scanln(&year)
	fmt("Podaj miesiac: (np:04)")
	var month string
    fmt.Scanln(&month)
	fmt("Podaj dzień: (np:01)")
	var day string
    fmt.Scanln(&day)

	var date string = year+"-"+month+"-"+day


}

func DisplayCurrentCurrenciesBuySell() {

}

func DisplayCurrentCurrenciesBuySellOnDay() {

}

func CurrencyConversion() {

}

func Settings() {

}

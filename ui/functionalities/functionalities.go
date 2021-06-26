package functionalities

import (
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
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

func DisplayCurrentCurrenciesBuySell(date string) {
	currencies := web.GetCurrencies(date)

	var rows [][]string

	for _, currency := range currencies.Items {
		val := strconv.FormatFloat(float64(currency.Value), 'f', 4, 64)
		sellValue := strconv.FormatFloat(float64(currency.Value)*float64(1.05), 'f', 4, 64)
		rows = append(rows, []string{
			currency.Name,
			val,
			sellValue,
			currency.Code,
		})
	}

	table.ShowMyPrettyTableBuyAndSell(rows)
}

func DisplayCurrentCurrenciesBuySellOnDay(date string) {
	currencies := web.GetCurrencies(date)

	var rows [][]string

	for _, currency := range currencies.Items {
		val := strconv.FormatFloat(float64(currency.Value), 'f', 4, 64)
		sellValue := strconv.FormatFloat(float64(currency.Value)*float64(1.05), 'f', 4, 64)
		rows = append(rows, []string{
			currency.Name,
			val,
			sellValue,
			currency.Code,
		})
	}

	table.ShowMyPrettyTableBuyAndSell(rows)
}

func CurrencyConversion(date string) [][]string {
	currencies := web.GetCurrencies(date)
	var amountOfMoney float64
	var rows [][]string
	var options []string
	fmt.Print(currencies)
	for _, currency := range currencies.Items {
		options = append(options, currency.Name)

	}

	if len(options) != 0 {
		prompt := promptui.Select{
			Label: "Wybierz opcję: ",
			Items: options,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return rows
		}

		fmt.Printf("Wybrana waluta to:" + result)
		fmt.Printf("Wprowadź ilość:")
		fmt.Scan(&amountOfMoney)

		for _, currency := range currencies.Items {
			val := strconv.FormatFloat(float64(currency.Value), 'f', 4, 64)
			multiplied_val := strconv.FormatFloat(float64(currency.Value)*float64(amountOfMoney), 'f', 4, 64)

			rows = append(rows, []string{
				currency.Name,
				multiplied_val,
				val,
				currency.Code,
			})
		}

	}
	return rows
}

func Settings() {

}

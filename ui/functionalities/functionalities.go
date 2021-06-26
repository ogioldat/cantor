package functionalities

import (
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/ogioldat/cantor/ui/table"
	"github.com/ogioldat/cantor/web"
	"github.com/pterm/pterm"
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

	table.ShowMyPrettyTable(rows, []string{"Waluta", "Cena", "Kod"})
}

func DisplayCurrentCurrenciesBuySell() {

}

func DisplayCurrentCurrenciesBuySellOnDay() {

}

func CurrencyConversion(date string) {
	currencies := web.GetCurrencies(date)
	var amountOfMoney float64
	var rows [][]string
	var options []string
	for _, currency := range currencies.Items {
		options = append(options, currency.Name)
	}

	if len(options) != 0 {
		prompt := promptui.Select{
			Label: "Wybierz walutę",
			Items: options,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Prompt failed %v\n", err)
			return
		}

		pterm.Info.Println("Wybrana waluta to " + pterm.LightMagenta(result))
		pterm.Info.Print("Wprowadź ilość (PLN): ")
		fmt.Scan(&amountOfMoney)

		for _, currency := range currencies.Items {
			if currency.Name == result {
				val := strconv.FormatFloat(float64(currency.Value), 'f', 4, 64)
				multiplied_val := strconv.FormatFloat(float64(currency.Value)*float64(amountOfMoney), 'f', 4, 64)

				rows = append(rows, []string{
					currency.Name,
					multiplied_val,
					val,
					currency.Code,
				})
				break
			}
		}
		table.ShowMyPrettyTable(rows, []string{"Waluta", "Wartość", "Cena", "Kod"})

	}
}

func Settings() {

}

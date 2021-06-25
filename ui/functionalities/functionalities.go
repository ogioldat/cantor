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

}

func DisplayCurrentCurrenciesBuySell() {

}

func DisplayCurrentCurrenciesBuySellOnDay() {

}

func CurrencyConversion() {

}

func Settings() {
	prompt := promptui.Select{
		Label: "Wybierz opcję: ",
		Items: []string{"Powrót"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	if result == "Powrót" {
		print("\033[H\033[2J")
	}
}

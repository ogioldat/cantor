package functionalities

import (
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/ogioldat/cantor/config"
	"github.com/ogioldat/cantor/ui/table"
	"github.com/ogioldat/cantor/utils"
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

func DisplayCurrentCurrenciesBuySell(date string) {
	currencies := web.GetCurrencies(date)
	cfg, _ := config.GetConfig()

	var rows [][]string

	for _, currency := range currencies.Items {
		val := strconv.FormatFloat(float64(currency.Value), 'f', 4, 64)
		sellValue := strconv.FormatFloat(float64(currency.Value)*float64(cfg.Provision), 'f', 4, 64)
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
	cfg, _ := config.GetConfig()

	var rows [][]string

	for _, currency := range currencies.Items {
		val := strconv.FormatFloat(float64(currency.Value), 'f', 4, 64)
		sellValue := strconv.FormatFloat(float64(currency.Value)*float64(cfg.Provision), 'f', 4, 64)
		rows = append(rows, []string{
			currency.Name,
			val,
			sellValue,
			currency.Code,
		})
	}

	table.ShowMyPrettyTableBuyAndSell(rows)
}

func CurrencyConversion(date string) {
	currencies := web.GetCurrencies(date)
	var amountOfMoney float64
	var rows [][]string
	var options []string
	for _, currency := range currencies.Items {
		options = append(options, currency.Name)
	}

	cfg, _ := config.GetConfig()

	if len(options) != 0 {
		prompt := promptui.Select{
			Label: "Wybierz walutę",
			Items: options,
		}

		_, result, err := prompt.Run()

		if err != nil {
			fmt.Printf("Wystąpił nieoczekiwany błąd %v\n", err)
			return
		}

		pterm.Info.Println("Wybrana waluta to " + pterm.LightMagenta(result))
		pterm.Info.Print("Wprowadź ilość (PLN): ")
		fmt.Scan(&amountOfMoney)

		for _, currency := range currencies.Items {
			if currency.Name == result {
				val := strconv.FormatFloat(float64(currency.Value), 'f', 4, 64)
				multiplied_val := strconv.FormatFloat(float64(currency.Value)*amountOfMoney*float64(cfg.Provision), 'f', 4, 64)

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
	cfg, _ := config.GetConfig()

	provision_percent := (cfg.Provision - 1) * 100
	provision_str := strconv.FormatFloat(provision_percent, 'f', 2, 64)

	provision_msg := "Ustaw prowizję (domyślnie " + provision_str + "% )"

	prompt := promptui.Select{
		Label: "Wybierz walutę",
		Items: []string{provision_msg, "Powrót"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Wystąpił nieoczekiwany błąd %v\n", err)
		return
	}

	if result == provision_msg {
		prompt := promptui.Prompt{
			Label:    "Number",
			Validate: utils.ValidateProvision,
		}

		result, _ := prompt.Run()
		fl, _ := strconv.ParseFloat(result, 64)
		config.SetProvision(1 + fl/100)
	}

	if result == "Powrót" {
		return
	}

}

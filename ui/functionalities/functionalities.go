package functionalities

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func DisplayCurrentCurrencies() {

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

package functionalities

<<<<<<< HEAD
import (
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/ogioldat/cantor/ui/table"
	"github.com/ogioldat/cantor/web"
)

func DisplayLatestCurrencies() {
	currencies := web.GetLatestCurrencies()
=======
func DisplayCurrentCurrencies() {
>>>>>>> aeba3ad8e91e2d1dd2265bc26d3d4c97c0519f0c

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

}

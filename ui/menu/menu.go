package menu

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/ogioldat/cantor/config"
	"github.com/ogioldat/cantor/ui/functionalities"
	"github.com/ogioldat/cantor/utils"
	"github.com/pterm/pterm"
)

var cfg, helpers = config.GetConfig()

func ShowHeader(headstring string) {
	ptermLogo, _ := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("KANTOR", pterm.NewStyle(pterm.FgLightCyan)),
		pterm.NewLettersFromStringWithStyle("GO", pterm.NewStyle(pterm.FgLightMagenta))).
		Srender()

	pterm.DefaultCenter.Print(ptermLogo)
	pterm.DefaultCenter.Print(pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Sprint(headstring))

}

func Showmenu() {

	var menuItems []string

	// fmt.Println(cfg.MenuOptions)

	for _, option := range cfg.MenuOptions {
		// fmt.Println(option.Name)
		menuItems = append(menuItems, option.Name)
	}

	fmt.Printf(menuItems[3])

	ShowHeader("MENU GŁÓWNE")
	prompt := promptui.Select{
		Label: "Wybierz opcję: ",
		Items: menuItems,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	if result == "Wyświetl aktualny kursów walut obcych" {
		ShowHeader("Wyświetl aktualny kursów walut obcych")
		functionalities.DisplayCurrencies("today")
	}
	if result == "Wyświetl kursów walut obcych z dnia wprowadzonego przez użytkownika" {
		ShowHeader("Wyświetl kursów walut obcych z dnia wprowadzonego przez użytkownika")
		pterm.DefaultHeader.Println("Wprowadź datę!")
		functionalities.DisplayCurrencies(utils.GetDate())
		// functionalities.DisplayCurrencies("2021-06-23")
	}
	if result == "Wyświetl aktuale kursy kupna i sprzedaży walut obcych" {
		ShowHeader("Wyświetl aktuale kursy kupna i sprzedaży walut obcych")
		functionalities.DisplayCurrentCurrenciesBuySell()
	}

	if result == "Przelicz waluty po aktualnej cenie bądź cenie z wybranego dnia" {
		ShowHeader("Przelicz waluty po aktualnej cenie bądź cenie z wybranego dnia")
		functionalities.CurrencyConversion()

	}
	if result == "Ustawienia" {
		ShowHeader("Ustawienia")
		functionalities.Settings()
	}
	if result == "Zakończ program" {
		print("\033[H\033[2J")
		os.Exit(3)
	}
}

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

var cfg, _ = config.GetConfig()

func ShowHeader(headstring string) {
	fmt.Println()
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
	ShowHeader("MENU GŁÓWNE")
	prompt := promptui.Select{
		Label: "Wybierz opcję",
		Items: []string{"Wyświetl aktualny kursów walut obcych",
			"Wyświetl kursów walut obcych z dnia wprowadzonego przez użytkownika",
			"Wyświetl aktuale kursy kupna i sprzedaży walut obcych",
			"Wyświetl kursy kupna i sprzedaży walut obcych z wybranego dnia",
			"Przelicz waluty po aktualnej cenie bądź cenie z wybranego dnia",
			"Ustawienia",
			"Zakończ program"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Wystąpił nieoczekiwany błąd %v\n", err)
		return
	}
	if result == "Wyświetl aktualny kursów walut obcych" {
		ShowHeader("Wyświetl aktualny kursów walut obcych")
		functionalities.DisplayCurrencies("last")
	}
	if result == "Wyświetl kursów walut obcych z dnia wprowadzonego przez użytkownika" {
		ShowHeader("Wyświetl kursów walut obcych z dnia wprowadzonego przez użytkownika")
		pterm.DefaultHeader.Println("Wprowadź datę!")
		functionalities.DisplayCurrencies(utils.GetDate())
	}
	if result == "Wyświetl aktuale kursy kupna i sprzedaży walut obcych" {
		ShowHeader("Wyświetl aktuale kursy kupna i sprzedaży walut obcych")
		functionalities.DisplayCurrentCurrenciesBuySell("last")
	}
	if result == "Wyświetl kursy kupna i sprzedaży walut obcych z wybranego dnia" {
		ShowHeader("Wyświetl kursy kupna i sprzedaży walut obcych z wybranego dnia")
		functionalities.DisplayCurrentCurrenciesBuySellOnDay(utils.GetDate())
	}

	if result == "Przelicz waluty po aktualnej cenie bądź cenie z wybranego dnia" {
		ShowHeader("Przelicz waluty po aktualnej cenie bądź cenie z wybranego dnia")

		prompt := promptui.Select{
			Label: "Wybierz akcję",
			Items: []string{"Użyj ostatniej daty", "Wprowadź datę"},
		}

		_, result, _ := prompt.Run()

		if result == "Wprowadź datę" {
			pterm.DefaultHeader.Println("Wprowadź datę!")
			functionalities.CurrencyConversion(utils.GetDate())
		} else {
			pterm.DefaultHeader.Println("Wprowadź datę!")
			functionalities.CurrencyConversion("last")
		}

	}
	if result == "Ustawienia" {
		ShowHeader("Ustawienia")
		functionalities.Settings()
		Showmenu()
	}
	if result == "Zakończ program" {
		print("\033[H\033[2J")
		os.Exit(3)
	}
}

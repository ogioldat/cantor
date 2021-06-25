package menu

import (
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/pterm/pterm"
)

func ShowHeader() {}

func Showmenu() {

	ptermLogo, _ := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("CANTOR", pterm.NewStyle(pterm.FgLightCyan)),
		pterm.NewLettersFromStringWithStyle("GO", pterm.NewStyle(pterm.FgLightMagenta))).
		Srender()

	pterm.DefaultCenter.Print(ptermLogo)
	pterm.DefaultCenter.Print(pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Sprint("MENU GŁÓWNE"))

	prompt := promptui.Select{
		Label: "Wybierz opcję: ",
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
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	if result == "Wyświetl aktualny kursów walut obcych" {

	}
	if result == "Wyświetl kursów walut obcych z dnia wprowadzonego przez użytkownika" {

	}
	if result == "Wyświetl aktuale kursy kupna i sprzedaży walut obcych" {

	}

	if result == "Wyświetl kursy kupna i sprzedaży walut obcych z wybranego dnia" {

	}
	if result == "Przelicz waluty po aktualnej cenie bądź cenie z wybranego dnia" {

	}
	if result == "Ustawienia" {

	}
	if result == "Zakończ program" {
		print("\033[H\033[2J")
		os.Exit(3)
	}
}

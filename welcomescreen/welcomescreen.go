package welcomescreen

import (
	"strconv"
	"time"

	"github.com/pterm/pterm"
)

const second = time.Second

// func main() {
// 	introScreen()
// 	clear()
// }

func introScreen() {
	pterm.DefaultCenter.Print(pterm.Green("Wydział Matematyki stosowanej Politechniki Śląskiej\n"))

	ptermLogo, _ := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("CANTOR", pterm.NewStyle(pterm.FgLightCyan)),
		pterm.NewLettersFromStringWithStyle("GO", pterm.NewStyle(pterm.FgLightMagenta))).
		Srender()

	pterm.DefaultCenter.Print(ptermLogo)

	pterm.DefaultCenter.Print(pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Sprint("Kantor walut stowony w języku GoLang"))

	pterm.Info.Println("Program stworzony przez " + pterm.LightMagenta("Michała Wieczorka") + " oraz " + pterm.LightMagenta("Tomasza Ogiołdę") +
		"\nw ramach Projektu zaliczeniowego z przedmiotu Programowanie 2 " +
		"\nwykładanego przez " + pterm.LightMagenta("dr.Sobotę.\n"))
	pterm.Println()

	introSpinner, _ := pterm.DefaultSpinner.WithRemoveWhenDone(true).Start("Uruchomienie prgramu nastąpi za 10s")
	time.Sleep(second)
	for i := 9; i > 0; i-- {
		if i > 1 {
			introSpinner.UpdateText("Uruchomienie prgramu nastąpi za " + strconv.Itoa(i) + " sekund...")
		} else {
			introSpinner.UpdateText("Uruchomienie prgramu nastąpi za " + strconv.Itoa(i) + " sekund...")
		}
		time.Sleep(second)
	}
	introSpinner.Stop()
}

func clear() {
	print("\033[H\033[2J")
}

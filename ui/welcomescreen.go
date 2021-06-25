package welcomescreen

import (
	"strconv"
	"time"

	"github.com/pterm/pterm"
)

const second = time.Second
const sleepTime = 5 // s

func GetMessage(t int) string {
	return "Uruchomienie prgramu nastąpi za " + strconv.Itoa(t) + "s"
}

func LaunchMessage(message string) {
	introSpinner, _ := pterm.DefaultSpinner.WithRemoveWhenDone(true).Start(GetMessage(sleepTime))

	for i := sleepTime; i > 0; i-- {
		introSpinner.UpdateText(GetMessage(i))

		time.Sleep(second)
	}
	introSpinner.Stop()
}

func IntroScreen() {
	pterm.DefaultCenter.Println(
		pterm.Green("Wydział Matematyki stosowanej Politechniki Śląskiej"))

	ptermLogo, _ := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("CANTOR", pterm.NewStyle(pterm.FgLightCyan)),
		pterm.NewLettersFromStringWithStyle("GO", pterm.NewStyle(pterm.FgLightMagenta))).
		Srender()

	pterm.DefaultCenter.Print(ptermLogo)

	pterm.DefaultCenter.Print(pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Sprint("Kantor walut stworzony w języku GoLang"))

	pterm.Info.Println(
		"Program stworzony przez " + pterm.LightMagenta("Michała Wieczorka") + " oraz " + pterm.LightMagenta("Tomasza Ogiołdę") +
			"\nw ramach Projektu zaliczeniowego z przedmiotu Programowanie 2 " +
			"\nwykładanego przez " + pterm.LightMagenta("dr.Sobotę.\n"))

	LaunchMessage("Uruchomienie prgramu nastąpi za " + strconv.Itoa(sleepTime) + "s")
	pterm.Println()

}

func Clear() {
	print("\033[H\033[2J")
}

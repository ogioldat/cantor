package table

import (
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/pterm/pterm"
)

func ShowMyPrettyTable(data [][]string, headers []string) {

	if len(data) == 0 {
		pterm.Info.Println("Brak danych! Podaj inną datę")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.SetBorder(false)
	table.AppendBulk(data)
	table.Render()
}

func ShowMyPrettyTableBuyAndSell(data [][]string) {

	if len(data) == 0 {
		pterm.Info.Println("Brak danych! Podaj inną datę")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Waluta", "Cena kupna", "Cena sprzedaży", "Kod"})
	table.SetBorder(false)
	table.AppendBulk(data)
	table.Render()
}

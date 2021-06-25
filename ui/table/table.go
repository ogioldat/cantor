package table

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func ShowMyPrettyTable(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Waluta", "Cena", "Kod"})
	table.SetBorder(false)
	table.AppendBulk(data)
	table.Render()
}

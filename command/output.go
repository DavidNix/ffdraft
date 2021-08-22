package command

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

type Presenter interface {
	Headers() []string
	Rows() [][]string
}

func PrintTable(pres Presenter) {
	var table = tablewriter.NewWriter(os.Stdout)
	table.SetHeader(pres.Headers())
	rows := pres.Rows()
	for i := range rows {
		table.Append(rows[i])
	}
	table.Render()
}

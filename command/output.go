package command

import (
	"os"

	"github.com/davidnix/ffdraft/players"
	"github.com/olekukonko/tablewriter"
)

var headers = []string{
	"Name",
	"Pos",
	"Team",
	"Floor",
	"Ceil",
	"Cost",
	"AAV",
	"Drop",
	"Tier",
	"Age",
	"Exp",
	"BYE",
	"ECR",
	"ADP",
	"Rank",
	"PRank",
	"VOR",
	"Risk",
}

type Rows []players.Player

func PrintTable(rows Rows) {
	var table = tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	for _, r := range rows {
		table.Append(r.Row())
	}
	table.Render()
}

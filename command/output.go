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
	"Age",
	"Exp",
	"BYE",
	"ECR",
	"Rank",
	"PRank",
	"Tier",
	"Drop",
	"VOR",
	"Risk",
	"ADP",
	"AAV",
	"Cost",
	"Ceil",
	"Floor",
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

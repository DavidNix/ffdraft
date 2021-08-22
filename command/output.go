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

func PrintTable(repo *players.Repo, rows Rows) {
	var table = tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	for _, player := range rows {
		table.Append(player.Row(repo.Position))
	}
	table.Render()
}

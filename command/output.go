package command

import (
	"github.com/davidnix/ffdraft/players"
	"github.com/olekukonko/tablewriter"
	"os"
)

var headers = []string{"ID", "Name", "Pos", "Team", "ECR", "Rank", "Pos Rank", "VOR", "Dropoff", "Floor", "Ceil", "AAV", "Cost", "Risk"}

type Rows []players.Player

func PrintTable(rows Rows) {
	var table = tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	for _, r := range rows {
		table.Append(r.Row())
	}
	table.Render()
}

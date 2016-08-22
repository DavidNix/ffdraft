package command

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

var headers = []string{"Name", "Pos", "Team", "ECR", "Rank", "Pos Rank", "VOR", "Dropoff", "Floor", "Ceil", "AAV", "Cost", "Risk"}

var table = tablewriter.NewWriter(os.Stdout)

type Rowable interface {
	Row() []string
}

type Rows []Rowable

func PrintTable(rows Rows) {
	table.SetHeader(headers)
	for _, r := range rows {
		table.Append(r.Row())
	}
	table.Render()
}

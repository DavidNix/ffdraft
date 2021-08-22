package presenter

import (
	"github.com/davidnix/ffdraft/players"
)

var startCols = []string{
	"Player",
	"Pos",
	"Team",
}

var endCols = []string{
	"Tier",
	"Rank",
	"PRank",
	"Drop",
	"StdDev",
	"VOR",
	"Risk",
	"Injury",
}

func buildRow(p players.Player, middleVals ...string) (row []string) {
	row = append(row,
		p.Name,
		p.Position,
		p.Team,
	)
	row = append(row, middleVals...)
	row = append(row,
		Int(p.Tier).String(),
		Int(p.OverallRank).String(),
		Int(p.PositionRank).String(),
		Float(p.Dropoff).String(),
		Float(p.StdDevPoints).String(),
		Float(p.Vor).String(),
		Float(p.Risk).String(),
		Injury(p.Injury).String(),
	)
	return row
}

type Players []players.Player

func (ps Players) Headers() []string {
	cols := append([]string{}, startCols...)
	cols = append(cols, "Floor", "Ceil")
	return append(cols, endCols...)
}

func (ps Players) Rows() (rows [][]string) {
	for _, p := range ps {
		row := buildRow(p,
			Float(p.Floor).String(),
			Float(p.Ceil).String(),
		)
		rows = append(rows, row)
	}
	return rows
}

type FloorPlayers []players.Player

func (fp FloorPlayers) Headers() []string {
	cols := append([]string{}, startCols...)
	cols = append(cols, "Floor", "FRank", "FVor")
	return append(cols, endCols...)
}

func (fp FloorPlayers) Rows() (rows [][]string) {
	for _, p := range fp {
		row := buildRow(p,
			Float(p.Floor).String(),
			Float(p.FloorRank).String(),
			Float(p.FloorVor).String(),
		)
		rows = append(rows, row)
	}
	return rows
}

type CeilPlayers []players.Player

func (cp CeilPlayers) Headers() []string {
	cols := append([]string{}, startCols...)
	cols = append(cols, "Ceil", "CRank", "CVor")
	return append(cols, endCols...)
}

func (cp CeilPlayers) Rows() (rows [][]string) {
	for _, p := range cp {
		row := buildRow(p,
			Float(p.Ceil).String(),
			Float(p.CeilRank).String(),
			Float(p.CeilVor).String(),
		)
		rows = append(rows, row)
	}
	return rows
}

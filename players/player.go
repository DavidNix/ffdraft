package players

import (
	"strings"
)

const (
	QB  = "QB"
	WR  = "WR"
	RB  = "RB"
	TE  = "TE"
	DST = "DST"
	K   = "K"
)

func OrderedPositions() []string {
	return []string{
		RB,
		WR,
		TE,
		QB,
		DST,
		K,
	}
}

func Positions() map[string]bool {
	pos := make(map[string]bool)
	for _, v := range OrderedPositions() {
		pos[v] = true
	}
	return pos
}

// Player of the NFL variety
type Player struct {
	Ceil     float64 `json:"ceiling"`
	CeilRank float64 `json:"ceiling_rank"`
	CeilVor  float64 `json:"ceiling_vor"`

	Floor     float64 `json:"floor"`
	FloorRank float64 `json:"floor_rank"`
	FloorVor  float64 `json:"floor_vor"`

	ADP          float64 `json:"adp"`
	Age          int     `json:"age"`
	Dropoff      float64 `json:"drop_off"`
	Exp          int     `json:"exp"`
	ID           int     `json:"id"`
	NameFirst    string  `json:"first_name"`
	NameLast     string  `json:"last_name"`
	OverallRank  int     `json:"rank"`
	Position     string  `json:"pos"`
	PositionRank int     `json:"pos_rank"`
	StdDevPoints float64 `json:"sd_pts"`
	Team         string  `json:"team"`
	Tier         int     `json:"tier"`
	Vor          float64 `json:"points_vor"`
	Risk         float64 `json:"risk"`
}

func (p Player) Name() string {
	return p.NameFirst + " " + p.NameLast
}

func (p Player) String() string {
	return strings.Join([]string{p.Name(), p.Position, p.Team}, " ")
}

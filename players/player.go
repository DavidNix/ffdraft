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

	Dropoff      float64 `json:"dropoff"`
	Injury       string  `json:"injury_status"`
	Name         string  `json:"player"`
	OverallRank  int     `json:"rank"`
	Position     string  `json:"pos"`
	PositionRank int     `json:"pos_rank"`
	Risk         float64 `json:"risk"`
	StdDevPoints float64 `json:"sd_pts"`
	Team         string  `json:"team"`
	Tier         int     `json:"tier"`
	Vor          float64 `json:"points_vor"`
}

func (p Player) String() string {
	return strings.Join([]string{p.Name, p.Position, p.Team}, " ")
}

package players

import (
	"fmt"
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

// Response from fantasyfootballanalytics.net in format Data.pointsTable -> array of players
type Response struct {
	Data struct {
		Players []Player `json:"pointsTable"`
	} `json:"Data"`
}

// Player of any type
type Player struct {
	ID                int     `json:"playerId"`
	Name              string  `json:"playername"`
	Position          string  `json:"position"`
	Team              string  `json:"team"`
	VOR               float64 `json:"vor"`
	Points            float64 `json:"points"`
	ECR               float64 `json:"overallECR"`
	OverallRank       float64 `json:"overallRank"`
	PositionRank      float64 `json:"positionRank"`
	TargetAuctionCost float64 `json:"cost"`
	Dropoff           float64 `json:"dropoff"`
	AAV               float64 `json:"auctionValue"`
	Ceil              float64 `json:"upper"`
	Floor             float64 `json:"lower"`
	Risk              float64 `json:"risk"`
}

func (p Player) Row() []string {
	return []string{
		p.Name,
		p.Position,
		p.Team,
		fmt.Sprint(p.ECR),
		fmt.Sprint(p.OverallRank),
		fmt.Sprint(p.PositionRank),
		fmt.Sprint(p.VOR),
		fmt.Sprint(p.Dropoff),
		fmt.Sprint(p.Floor),
		fmt.Sprint(p.Ceil),
		fmt.Sprint(p.AAV),
		fmt.Sprint(p.TargetAuctionCost),
		fmt.Sprint(p.Risk),
	}
}

package players

import (
	"fmt"
	"math"
	"reflect"
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

type currency float64

// Player of any type
type Player struct {
	ID                int      `json:"playerId"`
	Name              string   `json:"player"`
	Position          string   `json:"position"`
	Team              string   `json:"team"`
	Floor             float64  `json:"lower"`
	Ceil              float64  `json:"upper"`
	TargetAuctionCost currency `json:"cost"`
	AAV               currency `json:"auctionValue"`
	Dropoff           float64  `json:"dropoff"`
	Tier              int      `json:"tier"`
	Age               int      `json:"age"`
	Exp               int      `json:"exp"`
	ByeWeek           int      `json:"bye"`
	ECR               float64  `json:"overallECR"`
	ADP               float64  `json:"adp"`
	OverallRank       int      `json:"overallRank"`
	PositionRank      int      `json:"positionRank"`
	VOR               float64  `json:"vor"`
	Risk              float64  `json:"risk"`
}

func (p Player) Row() (row []string) {
	val := reflect.ValueOf(p)
	if p.ID == 0 {
		return make([]string, val.NumField()-1)
	}
	for i := 0; i < val.NumField(); i++ {
		if i == 0 {
			continue // skip ID
		}
		var strVal string
		switch v := val.Field(i).Interface().(type) {
		case string:
			strVal = v
		case int:
			strVal = fmt.Sprint(v)
		case float64:
			strVal = fmt.Sprintf("%.2f", v)
		case currency:
			c := float64(v)
			if c/math.Trunc(c)-1 == 0 {
				strVal = fmt.Sprintf("$ %2.0f", v)
			} else {
				strVal = fmt.Sprintf("$ %.2f", v)
			}

		default:
			strVal = "<error>"
		}
		row = append(row, strVal)
	}
	return row
}

func (p Player) ShortDesc() string {
	return strings.Join(p.Row()[:3], " ")
}

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
	Position          string   `json:"playerposition"`
	Team              string   `json:"team"`
	Age               int      `json:"age"`
	Exp               int      `json:"exp"`
	ByeWeek           int      `json:"bye"`
	ECR               float64  `json:"overallECR"`
	OverallRank       int      `json:"overallRank"`
	PositionRank      int      `json:"positionRank"`
	Tier              int      `json:"tier"`
	Dropoff           float64  `json:"dropoff"`
	VOR               float64  `json:"vor"`
	Risk              float64  `json:"risk"`
	ADP               float64  `json:"adp"`
	TargetAuctionCost currency `json:"cost"`
	AAV               currency `json:"auctionValue"`
	Ceil              float64  `json:"upper"`
	Floor             float64  `json:"lower"`
}

func (p Player) Row() []string {
	val := reflect.ValueOf(p)
	row := make([]string, val.NumField())
	for i := 0; i < val.NumField(); i++ {
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
		row[i] = strVal
	}
	return row
}

func (p Player) ShortDesc() string {
	return strings.Join(p.Row()[:3], " ")
}

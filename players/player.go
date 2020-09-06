package players

import (
    "fmt"
    "math"
    "reflect"
    "strconv"
    "strings"

    "github.com/fatih/color"
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

func (p Player) Row(draftPos int) []string {
    if p.ID == 0 {
        val := reflect.ValueOf(p)
        return make([]string, val.NumField()-1)
    }

    return []string{
        p.Name,
        p.Position,
        p.Team,
        formatFloat(p.Floor),
        formatFloat(p.Ceil),
        formatCurrency(p.TargetAuctionCost),
        formatCurrency(p.AAV),
        formatFloat(p.Dropoff),
        formatInt(p.Tier),
        formatInt(p.Age),
        formatInt(p.Exp),
        formatInt(p.ByeWeek),
        formatFloat(p.ECR),
        formatADP(float64(draftPos), p.ADP),
        formatInt(p.OverallRank),
        formatInt(p.PositionRank),
        formatFloat(p.VOR),
        formatFloat(p.Risk),
    }
}

func formatInt(val int) string {
    return strconv.Itoa(val)
}

func formatFloat(val float64) string {
    return fmt.Sprintf("%.2f", val)
}

func formatCurrency(val currency) string {
    c := float64(val)
    if c/math.Trunc(c)-1 == 0 {
        return fmt.Sprintf("$ %2.0f", c)
    }
    return fmt.Sprintf("$ %.2f", c)
}

func formatADP(draftPos, adp float64) string {
    switch {
    case adp == 0:
        // noop
    case draftPos-adp > 10:
        return color.RedString("%.2f", adp)
    case draftPos-adp > 0:
        return color.GreenString("%.2f", adp)
    }
    return fmt.Sprintf("%.2f", adp)
}

func (p Player) ShortDesc() string {
    return strings.Join(p.Row(0)[:3], " ")
}

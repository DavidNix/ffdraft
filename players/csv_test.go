package players

import (
    "testing"

    "github.com/stretchr/testify/require"
)

func TestLoadFromCSV(t *testing.T) {
    p, err := LoadFromCSV("../fixtures/ffa_customrankings2017-0.csv")
    require.NoError(t, err)
    require.Len(t, p, 2056)

    lm := p[2]
    require.Equal(t, 79607, lm.ID)
    require.Equal(t, "LeSean McCoy", lm.Name)
    require.Equal(t, "RB", lm.Position)
    require.Equal(t, "BUF", lm.Team)
    require.Equal(t, 29, lm.Age)
    require.Equal(t, 8, lm.Exp)
    require.Equal(t, 6, lm.ByeWeek)
    require.Equal(t, 6.2, lm.ECR)
    require.Equal(t, 3, lm.OverallRank)
    require.Equal(t, 3, lm.PositionRank)
    require.Equal(t, 3, lm.Tier)
    require.Equal(t, 5.68541795564838, lm.Dropoff)
    require.Equal(t, 95.7526358379133, lm.VOR)
    require.Equal(t, 5.27772032926968, lm.Risk)
    require.Equal(t, 5.7575, lm.ADP)
    require.Equal(t, 56.0, float64(lm.TargetAuctionCost))
    require.Equal(t, 50.035, float64(lm.AAV))
    require.Equal(t, 226.556182279423, lm.Ceil)
    require.Equal(t, 186.3, lm.Floor)
}

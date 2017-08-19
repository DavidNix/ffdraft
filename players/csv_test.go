package players

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadFromCSV(t *testing.T) {
	p, err := LoadFromCSV("../fixtures/ffa_customrankings2017-0.csv")
	require.NoError(t, err)
	require.Len(t, p, 2056)

	lm := p[2]
	assert.Equal(t, 79607, lm.ID)
	assert.Equal(t, "LeSean McCoy", lm.Name)
	assert.Equal(t, "RB", lm.Position)
	assert.Equal(t, "BUF", lm.Team)
	assert.Equal(t, 29, lm.Age)
	assert.Equal(t, 8, lm.Exp)
	assert.Equal(t, 6, lm.ByeWeek)
	assert.Equal(t, 6.2, lm.ECR)
	assert.Equal(t, 3, lm.OverallRank)
	assert.Equal(t, 3, lm.PositionRank)
	assert.Equal(t, 3, lm.Tier)
	assert.Equal(t, 5.68541795564838, lm.Dropoff)
	assert.Equal(t, 95.7526358379133, lm.VOR)
	assert.Equal(t, 5.27772032926968, lm.Risk)
	assert.Equal(t, 5.7575, lm.ADP)
	assert.Equal(t, 56.0, float64(lm.TargetAuctionCost))
	assert.Equal(t, 50.035, float64(lm.AAV))
	assert.Equal(t, 226.556182279423, lm.Ceil)
	assert.Equal(t, 186.3, lm.Floor)
}

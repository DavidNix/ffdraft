package players

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/fixtures/rankings.csv
var fixtureCSV []byte

func TestLoadFromCSV(t *testing.T) {
	p, err := LoadFromCSV(bytes.NewReader(fixtureCSV))
	require.NoError(t, err)
	require.Equal(t, 32, len(p))

	want := Player{
		Ceil:         386,
		CeilRank:     25,
		CeilVor:      80.5,
		Dropoff:      4.88,
		Floor:        327,
		FloorRank:    23,
		FloorVor:     74.4,
		Injury:       "",
		Name:         "Kyler Murray",
		OverallRank:  28,
		Position:     "QB",
		PositionRank: 3,
		Risk:         3.64,
		StdDevPoints: 21,
		Team:         "ARI",
		Tier:         2,
		Vor:          72.3,
	}

	require.Equal(t, want, p[2])
}

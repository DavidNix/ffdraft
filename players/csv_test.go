package players

import (
	"bytes"
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed testdata/fixtures/projections.csv
var fixtureCSV []byte

func TestLoadFromCSV(t *testing.T) {
	got, err := LoadFromCSV(bytes.NewReader(fixtureCSV))
	require.NoError(t, err)
	require.NotEmpty(t, got)

	want := Player{
		ADP:          64.43,
		Age:          24,
		Ceil:         382.811,
		CeilRank:     24,
		CeilVor:      76.5103405358627,
		Dropoff:      18.1745431233252,
		Exp:          1,
		Floor:        328.143,
		FloorRank:    29,
		FloorVor:     72.0303333333333,
		ID:           14056,
		NameFirst:    "Kyler",
		NameLast:     "Murray",
		OverallRank:  24,
		Position:     "QB",
		PositionRank: 4,
		StdDevPoints: 20.9414080892839,
		Team:         "ARI",
		Tier:         2,
		Vor:          74.0442784496819,
	}

	require.Equal(t, want, got[2])
}

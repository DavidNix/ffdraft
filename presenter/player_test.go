package presenter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayers_Headers(t *testing.T) {
	pres := Players{}
	h := pres.Headers()
	require.Len(t, h, 13)

	require.Contains(t, h, "Ceil")
	require.Contains(t, h, "Floor")
}

func TestPlayers_Rows(t *testing.T) {
	pres := make(Players, 2)
	require.Len(t, pres.Rows(), 2)
	require.Equal(t, len(pres.Headers()), len(pres.Rows()[0]))
}

func TestCeilPlayers_Headers(t *testing.T) {
	pres := CeilPlayers{}
	h := pres.Headers()
	require.Len(t, h, 14)

	require.Contains(t, h, "Ceil")
	require.NotContains(t, h, "Floor")
}

func TestCeilPlayers_Rows(t *testing.T) {
	pres := make(CeilPlayers, 2)
	require.Len(t, pres.Rows(), 2)
	require.Equal(t, len(pres.Headers()), len(pres.Rows()[0]))
}

func TestFloorPlayers_Headers(t *testing.T) {
	pres := FloorPlayers{}
	h := pres.Headers()

	require.Len(t, h, 14)

	require.Contains(t, h, "Floor")
	require.NotContains(t, h, "Ceil")
}

func TestFloorPlayers_Rows(t *testing.T) {
	pres := make(FloorPlayers, 2)
	require.Len(t, pres.Rows(), 2)
	require.Equal(t, len(pres.Headers()), len(pres.Rows()[0]))
}

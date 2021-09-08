package presenter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayers(t *testing.T) {
	pres := make(Players, 2)
	h := pres.Headers()
	require.Contains(t, h, "Ceil")
	require.Contains(t, h, "Floor")
}

func TestCeilPlayers(t *testing.T) {
	pres := make(CeilPlayers, 2)
	h := pres.Headers()
	require.Contains(t, h, "Ceil")
	require.NotContains(t, h, "Floor")
}

func TestFloorPlayers(t *testing.T) {
	pres := make(FloorPlayers, 2)
	h := pres.Headers()
	require.Contains(t, h, "Floor")
	require.NotContains(t, h, "Ceil")
}

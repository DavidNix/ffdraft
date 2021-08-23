package players

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBy_Sort(t *testing.T) {
	plys := []Player{
		{NameFirst: "Peyton", NameLast: "Manning", Floor: 1},
		{NameFirst: "Russel", NameLast: "Wilson", Floor: 2},
		{NameFirst: "Emmit", NameLast: "Smith", Floor: 4},
		{NameFirst: "Khalil", NameLast: "Mack", Floor: 10},
		{NameFirst: "David", NameLast: "Carr", Floor: 3},
	}

	floor := func(p1, p2 Player) bool {
		return p1.Floor > p2.Floor
	}
	By(floor).Sort(plys)

	require.Equal(t, plys[0].Name(), "Khalil Mack")
	require.Equal(t, plys[1].Name(), "Emmit Smith")
	require.Equal(t, plys[2].Name(), "David Carr")
	require.Equal(t, plys[3].Name(), "Russel Wilson")
	require.Equal(t, plys[4].Name(), "Peyton Manning")
}

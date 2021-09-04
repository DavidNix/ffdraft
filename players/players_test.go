package players

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayers_Filter(t *testing.T) {
	players := Players{
		{ID: 1},
		{ID: 4},
		{ID: 2},
		{ID: 3},
	}

	got := players.Filter(func(p Player) bool {
		return p.ID > 2
	})

	require.Equal(t, Players{{ID: 4}, {ID: 3}}, got)
}

func TestPlayers_GroupPosition(t *testing.T) {
	players := Players{
		{ID: 7, Position: "TE"},
		{ID: 1, Position: "DST"},
		{ID: 2, Position: "QB"},
		{ID: 3, Position: "K"},
		{ID: 5, Position: "RB"},
		{ID: 4, Position: "RB"},
		{ID: 9, Position: "RB"},
		{ID: 6, Position: "WR"},
	}

	got := players.GroupPosition(func(p1, p2 Player) bool {
		return p1.ID < p2.ID
	}, 2)

	want := Players{
		{},
		{ID: 4, Position: "RB"},
		{ID: 5, Position: "RB"},
		{},
		{ID: 6, Position: "WR"},
		{},
		{ID: 7, Position: "TE"},
		{},
		{ID: 2, Position: "QB"},
		{},
		{ID: 1, Position: "DST"},
		{},
		{ID: 3, Position: "K"},
	}

	require.Equal(t, want, got)
}

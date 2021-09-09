package players

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTeam_Sync(t *testing.T) {
	plyrs := Players{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	}
	team := &Team{Players: []TeamPlayer{{ID: 12345}}}
	team.Sync(plyrs)

	want := []TeamPlayer{{ID: 1}, {ID: 2}, {ID: 3}}
	require.Equal(t, want, team.Players)
}

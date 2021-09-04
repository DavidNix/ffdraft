package players

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var claimed = []Player{
	{ID: 1, NameFirst: "Antonio", NameLast: "Brown", Position: "WR"},
	{ID: 2, NameFirst: "Tony", NameLast: "Romo", Position: "QB"},
}

var avail = []Player{
	{ID: 3, NameFirst: "Jason", NameLast: "Witten", Position: "TE"},
	{ID: 5, NameFirst: "Joshua Smith", Position: "RB"},
	{ID: 6, NameFirst: "Tim", NameLast: "Tebow", Position: "filtered out"},
	{ID: 4, NameFirst: "Josh", NameLast: "Gordon", Position: "WR"},
}

var subject = &Repo{
	Claimed:   claimed,
	Available: avail,
}

func TestRepo_Find(t *testing.T) {
	var tests = []struct {
		search   string
		wantName string
	}{
		{"antonio", "Antonio Brown"},
		{"a brown", "Antonio Brown"},
		{"witten", "Jason Witten"},
		{"josh gordon", "Josh Gordon"},
	}
	for _, tt := range tests {
		player := subject.FindAll(tt.search)[0]

		require.Equal(t, player.Name(), tt.wantName)
	}

	count := len(subject.FindAll("josh"))

	require.Equal(t, count, 2)
}

func TestNewRepo(t *testing.T) {
	r := NewRepo(avail)

	tebow := r.FindAll("Tim Tebow")
	require.Equal(t, len(tebow), 0)

	witten := r.FindAll("Witten")
	require.Equal(t, len(witten), 1)
}

func TestRepo_SyncTeam(t *testing.T) {
	player := Player{ID: 99, NameFirst: "Troy", NameLast: "Ache-man", Position: "QB"}
	all := append(avail, player)
	r := NewRepo(all)
	r.Claimed = make(Players, 3)
	availCount := len(r.Available)

	team := Team{
		Players: Players{
			player,
		},
	}

	r.SyncTeam(&team)

	require.Len(t, r.Claimed, 1)
	require.Equal(t, team.Players, r.Claimed)
	require.Len(t, r.Available, availCount-1)
}

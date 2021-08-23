package players

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var drafted = []Player{
	{NameFirst: "Antonio", NameLast: "Brown", Position: "WR"},
	{NameFirst: "Tony", NameLast: "Romo", Position: "QB"},
}

var unDrafted = []Player{
	{NameFirst: "Jason", NameLast: "Witten", Position: "TE"},
	{NameFirst: "Josh", NameLast: "Gordon", Position: "WR"},
	{NameFirst: "Joshua Smith", Position: "RB"},
	{NameFirst: "Tim", NameLast: "Tebow", Position: "Baseball"},
}

var subject = &Repo{
	Drafted:   drafted,
	UnDrafted: unDrafted,
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
	r := NewRepo(unDrafted)

	tebow := r.FindAll("Tim Tebow")
	require.Equal(t, len(tebow), 0)

	witten := r.FindAll("Witten")
	require.Equal(t, len(witten), 1)
}

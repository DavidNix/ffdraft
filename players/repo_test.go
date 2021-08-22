package players

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var drafted = []Player{
	{Name: "Antonio Brown", Position: "WR"},
	{Name: "Tony Romo", Position: "QB"},
}

var unDrafted = []Player{
	{Name: "Jason Witten", Position: "TE"},
	{Name: "Josh Gordon", Position: "WR"},
	{Name: "Joshua Smith", Position: "RB"},
	{Name: "Tim Tebow", Position: "Baseball"},
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

		require.Equal(t, player.Name, tt.wantName)
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

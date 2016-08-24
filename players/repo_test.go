package players

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var drafted []Player = []Player{
	{ID: 1, Name: "Antonio Brown", Position: "WR"},
	{ID: 2, Name: "Tony Romo", Position: "QB"},
}

var unDrafted []Player = []Player{
	{ID: 3, Name: "Jason Witten", Position: "TE"},
	{ID: 4, Name: "Josh Gordon", Position: "WR"},
	{ID: 5, Name: "Joshua Smith", Position: "RB"},
	{ID: 99, Name: "Tim Tebow", Position: "Baseball"},
}

var subject *Repo = &Repo{
	Drafted:   drafted,
	UnDrafted: unDrafted,
}

func TestRepo_Find(t *testing.T) {
	var tests = []struct {
		search string
		result int
	}{
		{"antonio", 1},
		{"a brown", 1},
		{"witten", 3},
		{"josh gordon", 4},
	}
	for _, test := range tests {
		player := subject.Find(test.search)[0]

		assert.Equal(t, player.ID, test.result)
	}

	count := len(subject.Find("josh"))

	assert.Equal(t, count, 2)
}

func TestNewRepo(t *testing.T) {
	r := NewRepo(unDrafted)

	tebow := r.Find("Tim Tebow")
	assert.Equal(t, len(tebow), 0)

	witten := r.Find("Witten")
	assert.Equal(t, len(witten), 1)
}

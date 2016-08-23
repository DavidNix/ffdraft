package players

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var drafted []Player = []Player{
	{ID: 1, Name: "Antonio Brown"},
	{ID: 2, Name: "Tony Romo"},
}

var unDrafted []Player = []Player{
	{ID: 3, Name: "Jason Witten"},
	{ID: 4, Name: "Josh Gordon"},
	{ID: 5, Name: "Joshua Smith"},
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

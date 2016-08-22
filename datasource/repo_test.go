package datasource

import (
	"github.com/davidnix/ffdraft/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

var drafted []models.Player = []models.Player{
	models.Player{ID: 1, Name: "Antonio Brown"},
	models.Player{ID: 2, Name: "Tony Romo"},
}

var unDrafted []models.Player = []models.Player{
	models.Player{ID: 3, Name: "Jason Witten"},
	models.Player{ID: 4, Name: "Josh Gordon"},
	models.Player{ID: 5, Name: "Joshua Smith"},
}

var subject *Repo = &Repo{
	drafted:   drafted,
	unDrafted: unDrafted,
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

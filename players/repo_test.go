package players

import (
    "testing"

    "github.com/stretchr/testify/require"
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
        player := subject.FindAll(test.search)[0]

        require.Equal(t, player.ID, test.result)
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

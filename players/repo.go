package players

import (
	"github.com/renstrom/fuzzysearch/fuzzy"
)

type Repo struct {
	Drafted   []Player
	UnDrafted []Player
}

func NewRepo(players []Player) *Repo {
	pos := Positions()
	f := func(p Player) bool {
		return pos[p.Position]
	}
	return &Repo{
		Drafted:   []Player{},
		UnDrafted: filter(players, f),
	}
}

func (r *Repo) Find(name string) []Player {
	all := []Player{}
	all = append(all, r.Drafted...)
	all = append(all, r.UnDrafted...)
	found := []Player{}
	for _, player := range all {
		if fuzzy.MatchFold(name, player.Name) {
			found = append(found, player)
		}
	}
	return found
}

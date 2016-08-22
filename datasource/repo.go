package datasource

import (
	"github.com/davidnix/ffdraft/models"
	"github.com/renstrom/fuzzysearch/fuzzy"
)

type Repo struct {
	drafted   []models.Player
	unDrafted []models.Player
}

func New(players []models.Player) *Repo {
	return &Repo{
		[]models.Player{},
		players,
	}
}

func (r *Repo) Find(name string) []models.Player {
	all := append(r.unDrafted, r.drafted...)
	found := []models.Player{}
	for _, player := range all {
		if fuzzy.MatchFold(name, player.Name) {
			found = append(found, player)
		}
	}
	return found
}

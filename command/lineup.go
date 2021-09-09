package command

import (
	"github.com/davidnix/ffdraft/players"
	"github.com/davidnix/ffdraft/presenter"
	"github.com/fatih/color"
)

func Lineup(repo *players.Repo, rankByFloor bool) {
	var sorted players.SortBy
	if rankByFloor {
		sorted = func(p1, p2 players.Player) bool { return p1.Floor > p2.Floor }
	} else {
		sorted = func(p1, p2 players.Player) bool { return p1.Ceil > p2.Ceil }
	}
	plyrs := repo.Claimed.GroupPosition(sorted, 100)
	if len(plyrs) == 0 {
		color.Yellow("No players on your team yet")
		return
	}
	PrintTable(presenter.Players(plyrs))
}

package command

import (
	"github.com/davidnix/ffdraft/players"
	"github.com/davidnix/ffdraft/presenter"
	"github.com/fatih/color"
)

func Lineup(repo *players.Repo) {
	sorted := func(p1, p2 players.Player) bool { return p1.Floor > p2.Floor }
	plyrs := repo.Claimed.GroupPosition(sorted, 100)
	if len(plyrs) == 0 {
		color.Yellow("No players on your team yet")
		return
	}
	PrintTable(presenter.Players(plyrs))
}

func LineupFloor(repo *players.Repo) {
	floor := players.SortBy(func(p1, p2 players.Player) bool {
		return p1.Floor > p2.Floor
	})
	floor.Sort(repo.Claimed)
	PrintTable(presenter.Players(repo.Claimed))
}

func LineupCeil(repo *players.Repo) {
	floor := players.SortBy(func(p1, p2 players.Player) bool {
		return p1.Ceil > p2.Ceil
	})
	floor.Sort(repo.Claimed)
	PrintTable(presenter.Players(repo.Claimed))
}

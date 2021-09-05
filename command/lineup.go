package command

import (
	"github.com/davidnix/ffdraft/players"
	"github.com/davidnix/ffdraft/presenter"
)

func Lineup(team *players.Team, all players.Players, rankByFloor bool) error {
	for i := range team.Players {
		p := team.Players[i]
		found := all.Filter(func(match players.Player) bool {
			return p.ID == match.ID
		})
		if len(found) >= 1 {
			team.Players[i] = found[0]
		} else {
			team.Players[i] = players.Player{
				ID:        p.ID,
				NameFirst: p.NameFirst,
				NameLast:  p.NameLast,
				Position:  p.Position,
				Team:      p.Team,
			}
		}
	}
	var sorted players.SortBy
	if rankByFloor {
		sorted = func(p1, p2 players.Player) bool { return p1.Floor > p2.Floor }
	} else {
		sorted = func(p1, p2 players.Player) bool { return p1.Ceil > p2.Ceil }
	}
	PrintTable(presenter.Players(team.Players.GroupPosition(sorted, 100)))
	return nil
}

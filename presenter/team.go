package presenter

import "github.com/davidnix/ffdraft/players"

type Team players.Players

func (team Team) Headers() []string {
	return []string{
		"Player",
		"Pos",
		"Team",
		"Age",
		"Exp",
	}
}

func (team Team) Rows() (rows [][]string) {
	for _, p := range team {
		if p.ID == 0 {
			// blank row
			rows = append(rows, make([]string, len(team.Headers())))
		} else {
			row := []string{
				p.Name(),
				p.Position,
				p.Team,
				Int(p.Age).String(),
				Int(p.Exp).String(),
			}
			rows = append(rows, row)
		}
	}
	return rows
}

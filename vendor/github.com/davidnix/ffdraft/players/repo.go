package players

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

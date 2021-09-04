package players

type Repo struct {
	Claimed   []Player
	Position  int
	Available []Player
}

func NewRepo(players []Player) *Repo {
	pos := Positions()
	f := func(p Player) bool {
		return pos[p.Position]
	}
	return &Repo{
		Claimed:   []Player{},
		Available: filter(players, f),
	}
}

package players

type Repo struct {
	Claimed   Players
	Position  int
	Available Players
}

func NewRepo(players Players) *Repo {
	pos := Positions()
	f := func(p Player) bool {
		return pos[p.Position]
	}
	return &Repo{
		Claimed:   Players{},
		Available: players.Filter(f),
	}
}

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

func (r *Repo) SyncTeam(t *Team) {
	r.Claimed = t.Players

	r.Available = r.Available.Filter(func(p Player) bool {
		for _, claimed := range r.Claimed {
			if claimed.ID != p.ID {
				return true
			}
		}
		return false
	})
}

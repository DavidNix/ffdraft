package players

type TeamPlayer struct {
	ID int
}

type Team struct {
	Name    string
	Players []TeamPlayer
}

func (t *Team) Sync(plyrs Players) {
	t.Players = nil
	for _, p := range plyrs {
		t.Players = append(t.Players, TeamPlayer{ID: p.ID})
	}
}

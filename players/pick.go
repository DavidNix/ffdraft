package players

import "github.com/pkg/errors"

func (r *Repo) Pick(p Player) error {
	for i, player := range r.Available {
		if p == player {
			r.Claimed = append(r.Claimed, p)
			r.Available = append(r.Available[:i], r.Available[i+1:]...)
			r.Position++
			return nil
		}
	}
	return errors.Errorf("pick: could not find player %s", p)
}

func (r *Repo) UnPick(p Player) error {
	for i, player := range r.Claimed {
		if p == player {
			r.Available = append(r.Available, p)
			r.Claimed = append(r.Claimed[:i], r.Claimed[i+1:]...)
			r.Position--
			return nil
		}
	}
	return errors.Errorf("unpick: could not find player %s", p)
}

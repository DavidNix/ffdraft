package players

import "github.com/pkg/errors"

func (r *Repo) Keep(p Player) error {
	for i, player := range r.Available {
		if p == player {
			r.Claimed = append(r.Claimed, p)
			r.Available = append(r.Available[:i], r.Available[i+1:]...)
			return nil
		}
	}
	return errors.Errorf("Keep: could not find player %s", p)
}

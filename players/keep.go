package players

import "github.com/pkg/errors"

func (r *Repo) Keep(p Player) error {
	for i, player := range r.UnDrafted {
		if p == player {
			r.Drafted = append(r.Drafted, p)
			r.UnDrafted = append(r.UnDrafted[:i], r.UnDrafted[i+1:]...)
			return nil
		}
	}
	return errors.Errorf("Keep: could not find player %s", p)
}

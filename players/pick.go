package players

import "github.com/pkg/errors"

func (r *Repo) Pick(p Player) error {
	for i, player := range r.UnDrafted {
		if p == player {
			r.Drafted = append(r.Drafted, p)
			r.UnDrafted = append(r.UnDrafted[:i], r.UnDrafted[i+1:]...)
			r.Position++
			return nil
		}
	}
	return errors.Errorf("pick: could not find player %s", p)
}

func (r *Repo) UnPick(p Player) error {
	for i, player := range r.Drafted {
		if p == player {
			r.UnDrafted = append(r.UnDrafted, p)
			r.Drafted = append(r.Drafted[:i], r.Drafted[i+1:]...)
			r.Position--
			return nil
		}
	}
	return errors.Errorf("unpick: could not find player %s", p)
}

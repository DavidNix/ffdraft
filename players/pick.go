package players

import "errors"

func (r *Repo) Pick(p Player) error {
	for i, player := range r.UnDrafted {
		if p == player {
			r.Drafted = append(r.Drafted, p)
			r.UnDrafted = append(r.UnDrafted[:i], r.UnDrafted[i+1:]...)
			r.Position++
			return nil
		}
	}
	return errors.New("pick: could not find player " + p.ShortDesc())
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
	return errors.New("unpick: could not find player " + p.ShortDesc())
}

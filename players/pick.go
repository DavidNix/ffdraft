package players

import "errors"

func (r *Repo) Pick(p Player) error {
    for i, player := range r.UnDrafted {
        if p.ID == player.ID {
            r.Drafted = append(r.Drafted, p)
            r.UnDrafted = append(r.UnDrafted[:i], r.UnDrafted[i+1:]...)
            r.Position++
            return nil
        }
    }
    return errors.New("Pick: Could not find player " + p.ShortDesc())
}

func (r *Repo) UnPick(p Player) error {
    for i, player := range r.Drafted {
        if p.ID == player.ID {
            r.UnDrafted = append(r.UnDrafted, p)
            r.Drafted = append(r.Drafted[:i], r.Drafted[i+1:]...)
            r.Position--
            return nil
        }
    }
    return errors.New("UnPick: Could not find player " + p.ShortDesc())
}

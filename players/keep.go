package players

import "errors"

func (r *Repo) Keep(p Player) error {
    for i, player := range r.UnDrafted {
        if p.ID == player.ID {
            r.Drafted = append(r.Drafted, p)
            r.UnDrafted = append(r.UnDrafted[:i], r.UnDrafted[i+1:]...)
            return nil
        }
    }
    return errors.New("Keep: could not find player " + p.ShortDesc())
}

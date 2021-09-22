package players

import "strings"

func (r *Repo) Floor() []Player {
	floor := func(p1, p2 Player) bool {
		return p1.Floor > p2.Floor
	}
	return r.Available.GroupPosition(floor, 5)
}

func (r *Repo) FloorByPos(pos string, limit int) []Player {
	plyrs := r.Available.Filter(func(p Player) bool {
		return strings.ToUpper(pos) == strings.ToUpper(p.Position) && p.Floor > 0
	})
	floor := func(p1, p2 Player) bool {
		return p1.Floor > p2.Floor
	}
	SortBy(floor).Sort(plyrs)
	return plyrs[:min(limit, len(plyrs))]
}

func (r *Repo) Ceil() []Player {
	ceil := func(p1, p2 Player) bool {
		return p1.Ceil > p2.Ceil
	}
	return r.Available.GroupPosition(ceil, 5)
}

func (r *Repo) CeilByPos(pos string, limit int) []Player {
	plyrs := r.Available.Filter(func(p Player) bool {
		return strings.ToUpper(pos) == strings.ToUpper(p.Position) && p.Ceil > 0
	})
	ceil := func(p1, p2 Player) bool {
		return p1.Ceil > p2.Ceil
	}
	SortBy(ceil).Sort(plyrs)
	return plyrs[:min(limit, len(plyrs))]
}

func (r *Repo) Team(name string) []Player {
	floor := func(p1, p2 Player) bool {
		return p1.Floor > p2.Floor
	}
	grouped := r.Available.GroupPosition(floor, 1000)
	return grouped.Filter(func(p Player) bool {
		return strings.ToUpper(name) == strings.ToUpper(p.Team) || p.Team == ""
	})
}

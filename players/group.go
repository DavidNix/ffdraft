package players

import "strings"

func (r *Repo) Floor() []Player {
	floor := func(p1, p2 Player) bool {
		return p1.Floor > p2.Floor
	}
	return r.group(floor, 5)
}

func (r *Repo) FloorByPos(pos string) []Player {
    plyrs := filter(r.UnDrafted, func (p Player) bool {
        return strings.ToUpper(pos) == strings.ToUpper(p.Position)
    })
    floor := func(p1, p2 Player) bool {
        return p1.Floor > p2.Floor
    }
    By(floor).Sort(plyrs)
    return plyrs[:min(30, len(plyrs))]
}

func (r *Repo) Ceil() []Player {
	ceil := func(p1, p2 Player) bool {
		return p1.Ceil > p2.Ceil
	}
	return r.group(ceil, 5)
}

func (r *Repo) CeilByPos(pos string) []Player {
    plyrs := filter(r.UnDrafted, func (p Player) bool {
        return strings.ToUpper(pos) == strings.ToUpper(p.Position)
    })
    ceil := func(p1, p2 Player) bool {
        return p1.Ceil > p2.Ceil
    }
    By(ceil).Sort(plyrs)
    return plyrs[:min(30, len(plyrs))]
}

func (r *Repo) Team(name string) []Player {
	floor := func(p1, p2 Player) bool {
		return p1.Floor > p2.Floor
	}
	grouped := r.group(floor, 1000)
	return filter(grouped, func(p Player) bool {
		return strings.ToUpper(name) == strings.ToUpper(p.Team) || p.Team == ""
	})
}

func (r *Repo) group(sortFunc By, max int) []Player {
	results := []Player{}
	for _, pos := range OrderedPositions() {
		undrafted := filter(r.UnDrafted, func(p Player) bool {
			return p.Position == pos
		})
		By(sortFunc).Sort(undrafted)
		l := limit(undrafted, max)
		results = append(results, Player{})
		results = append(results, l...)
	}
	return results
}

func limit(plyrs []Player, max int) []Player {
	if len(plyrs) == 0 {
		return plyrs
	}
	index := min(max, len(plyrs))
	return plyrs[:index]
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

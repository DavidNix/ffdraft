package players

func (r *Repo) Floor() []Player {
	floor := func(p1, p2 Player) bool {
		return p1.Floor > p2.Floor
	}
	return r.group(floor)
}

func (r *Repo) Ceil() []Player {
    ceil := func(p1, p2 Player) bool {
        return p1.Ceil > p2.Ceil
    }
    return r.group(ceil)
}

func (r *Repo) group(sortFunc By) []Player {
	results := []Player{}
	for _, pos := range OrderedPositions() {
		undrafted := filter(r.UnDrafted, func(p Player) bool {
			return p.Position == pos
		})
		By(sortFunc).Sort(undrafted)
		limited := limit(pos, undrafted)
		results = append(results, Player{})
		results = append(results, limited...)
	}
	return results
}

func limit(pos string, plyrs []Player) []Player {
	if len(plyrs) == 0 {
		return plyrs
	}
	index := min(3, len(plyrs))
	switch pos {
	case QB, WR, RB, TE, DST:
		return plyrs[:index]
	case K:
		return plyrs[:1]
	default:
		return plyrs
	}
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

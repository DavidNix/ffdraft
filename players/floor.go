package players

func (r *Repo) Floor() []Player {
	results := []Player{}
	for _, pos := range OrderedPositions() {
		undrafted := filter(r.UnDrafted, func(p Player) bool {
			return p.Position == pos
		})
		floor := func(p1, p2 Player) bool {
			return p1.Floor > p2.Floor
		}
		By(floor).Sort(undrafted)
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

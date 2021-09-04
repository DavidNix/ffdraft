package players

type Players []Player

type FilterFn func(p Player) bool

func (ps Players) Filter(f FilterFn) Players {
	var filtered []Player
	for _, player := range ps {
		if f(player) {
			filtered = append(filtered, player)
		}
	}
	return filtered
}

func (ps Players) GroupPosition(sort By, max int) (results Players) {
	for _, pos := range OrderedPositions() {
		players := ps.Filter(func(p Player) bool {
			return p.Position == pos
		})
		sort.Sort(players)
		l := limit(players, max)
		if len(l) == 0 {
			continue
		}
		results = append(results, Player{})
		results = append(results, l...)
	}
	return results
}

func limit(ps Players, max int) Players {
	if len(ps) == 0 {
		return ps
	}
	index := min(max, len(ps))
	return ps[:index]
}

func min(i, j int) int {
	if i < j {
		return i
	} else {
		return j
	}
}

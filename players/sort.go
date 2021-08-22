package players

import "sort"

type By func(p1, p2 Player) bool

type playerSorter struct {
	arr []Player
	by  By
}

func (by By) Sort(pls []Player) {
	ps := &playerSorter{pls, by}
	sort.Sort(ps)
}

func (ps *playerSorter) Swap(i, j int) {
	ps.arr[i], ps.arr[j] = ps.arr[j], ps.arr[i]
}

func (ps *playerSorter) Len() int {
	return len(ps.arr)
}

func (ps *playerSorter) Less(i, j int) bool {
	return ps.by(ps.arr[i], ps.arr[j])
}

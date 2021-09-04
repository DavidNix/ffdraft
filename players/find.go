package players

import "github.com/lithammer/fuzzysearch/fuzzy"

func (r *Repo) FindAll(name string) []Player {
	all := []Player{}
	all = append(all, r.Claimed...)
	all = append(all, r.Available...)
	return find(name, all)
}

func (r *Repo) FindAvailable(name string) []Player {
	return find(name, r.Available)
}

func (r *Repo) FindUnavailable(name string) []Player {
	return find(name, r.Claimed)
}

func find(name string, plyrs []Player) []Player {
	found := []Player{}
	for _, p := range plyrs {
		if fuzzy.MatchFold(name, p.Name()) {
			found = append(found, p)
		}
	}
	return found
}

package players

import "github.com/lithammer/fuzzysearch/fuzzy"

func (r *Repo) FindAll(name string) (all Players) {
	all = append(all, r.Claimed...)
	all = append(all, r.Available...)
	return find(name, all)
}

func (r *Repo) FindAvailable(name string) Players {
	return find(name, r.Available)
}

func (r *Repo) FindUnavailable(name string) Players {
	return find(name, r.Claimed)
}

func find(name string, plyrs Players) (found Players) {
	for _, p := range plyrs {
		if fuzzy.MatchFold(name, p.Name()) {
			found = append(found, p)
		}
	}
	return found
}

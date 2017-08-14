package players

import "github.com/renstrom/fuzzysearch/fuzzy"

func (r *Repo) FindAll(name string) []Player {
	all := []Player{}
	all = append(all, r.Drafted...)
	all = append(all, r.UnDrafted...)
	return find(name, all)
}

func (r *Repo) FindAvailable(name string) []Player {
	return find(name, r.UnDrafted)
}

func (r *Repo) FindUnavailable(name string) []Player {
	return find(name, r.Drafted)
}

func find(name string, plyrs []Player) []Player {
	found := []Player{}
	for _, p := range plyrs {
		if fuzzy.MatchFold(name, p.Name) {
			found = append(found, p)
		}
	}
	return found
}

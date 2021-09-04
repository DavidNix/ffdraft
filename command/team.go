package command

import (
	"log"
	"strings"

	"github.com/davidnix/ffdraft/players"
	"github.com/fatih/color"
)

func TeamAdd(r *players.Repo, t *players.Team, args []string) {
	name := strings.Join(args, " ")
	p, err := choose(r.FindAvailable(name))
	if err != nil {
		log.Println(err)
		return
	}

	if err := r.Pick(p); err != nil {
		log.Println(err)
		return
	}
	t.Players = r.Claimed
	color.Green("%s added to team %s", p, t.Name)
}

func TeamRemove(r *players.Repo, t *players.Team, args []string) {
	name := strings.Join(args, " ")
	p, err := choose(r.FindAvailable(name))
	if err != nil {
		log.Println(err)
		return
	}

	if err := r.UnPick(p); err != nil {
		log.Println(err)
		return
	}
	t.Players = r.Claimed
	color.Red("%s removed from team %s", p, t.Name)
}

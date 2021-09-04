package command

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/davidnix/ffdraft/players"
	"github.com/fatih/color"
)

var (
	errInvalidChoice  = errors.New("invalid choice; no player selected")
	errPlayerNotFound = errors.New("no players found")
)

func Pick(r *players.Repo, args []string) {
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
	pDesc := color.CyanString(p.String())
	log.Println(pDesc, "was picked.", len(r.Available), "available players remaining.")
}

func UnPick(r *players.Repo, args []string) {
	name := strings.Join(args, " ")
	p, err := choose(r.FindUnavailable(name))
	if err != nil {
		log.Println(err)
		return
	}

	if err := r.UnPick(p); err != nil {
		log.Println(err)
		return
	}
	pDesc := color.GreenString(p.String())
	log.Println(pDesc, "is now available.", len(r.Available), "available players remaining.")
}

func Keep(r *players.Repo, args []string) {
	name := strings.Join(args, " ")
	p, err := choose(r.FindAvailable(name))
	if err != nil {
		log.Println(err)
		return
	}
	if err := r.Keep(p); err != nil {
		log.Println(err)
		return
	}
	pDesc := color.MagentaString(p.String())
	log.Println(pDesc, "was kept.", len(r.Available), "available players remaining.")
}

func choose(choices []players.Player) (p players.Player, _ error) {
	if len(choices) == 0 {
		return p, errPlayerNotFound
	}

	for i, c := range choices {
		log.Printf("%v: %s", i+1, c.String()+"\n")
	}

	log.Print("Choose player:")
	in, err := GetInput()
	if err != nil {
		return p, err
	}
	selection, err := strconv.ParseInt(in, 0, 0)
	if err != nil {
		return p, errInvalidChoice
	}
	index := int(selection - 1)
	if index < 0 || index >= len(choices) {
		return p, errInvalidChoice
	}
	return choices[index], nil
}

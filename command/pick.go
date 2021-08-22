package command

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/davidnix/ffdraft/players"
	"github.com/fatih/color"
)

var invalidErr = errors.New("Invalid choice. No player selected.")

func Pick(r *players.Repo, args []string) {
	name := strings.Join(args, " ")

	p, err := choose(r.FindAvailable(name))
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := r.Pick(p); err != nil {
		fmt.Println(err)
		return
	}
	pDesc := color.CyanString(p.String())
	fmt.Println(pDesc, "was picked.", len(r.UnDrafted), "available players remaining.")
}

func UnPick(r *players.Repo, args []string) {
	name := strings.Join(args, " ")
	p, err := choose(r.FindUnavailable(name))
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := r.UnPick(p); err != nil {
		fmt.Println(err)
		return
	}
	pDesc := color.GreenString(p.String())
	fmt.Println(pDesc, "is now available.", len(r.UnDrafted), "available players remaining.")
}

func Keep(r *players.Repo, args []string) {
	name := strings.Join(args, " ")
	p, err := choose(r.FindAvailable(name))
	if err != nil {
		fmt.Println(err)
		return
	}
	if err := r.Keep(p); err != nil {
		fmt.Println(err)
		return
	}
	pDesc := color.MagentaString(p.String())
	fmt.Println(pDesc, "was kept.", len(r.UnDrafted), "available players remaining.")
}

func choose(choices []players.Player) (players.Player, error) {
	for i, p := range choices {
		fmt.Printf("%v: %s", i+1, p.String()+"\n")
	}

	fmt.Print("Choose player:")
	selection, err := strconv.ParseInt(GetInput('\n'), 0, 0)
	if err != nil {
		return players.Player{}, invalidErr
	}
	index := int(selection - 1)
	if index < 0 || index >= len(choices) {
		return players.Player{}, invalidErr
	}
	return choices[index], nil
}

package command

import (
	"fmt"
	"github.com/davidnix/ffdraft/players"
	"strconv"
	"strings"
	"errors"
)

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
	fmt.Println(p.ShortDesc(), "was picked.", len(r.UnDrafted), "available players remaining.")
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
	fmt.Println(p.ShortDesc(), "is now available.", len(r.UnDrafted), "available players remaining.")
}

func choose(choices []players.Player) (players.Player, error) {
	for i, p := range choices {
		fmt.Printf("%v: %s", i+1, p.ShortDesc() + "\n")
	}

	fmt.Print("Choose player:")
	selection, err := strconv.ParseInt(GetInput('\n'), 0, 0)
	if err != nil {
		return players.Player{}, err
	}
	index := int(selection - 1)
	if index < 0 || index >= len(choices) {
		return players.Player{}, errors.New("Invalid choice. No player selected.")
	}
	return choices[index], nil
}
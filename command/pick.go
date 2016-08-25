package command

import (
	"fmt"
	"github.com/davidnix/ffdraft/players"
	"strconv"
	"strings"
)

func Pick(r *players.Repo, args []string) {
	name := strings.Join(args, " ")
	choices := r.FindAvailable(name)
	for i, p := range choices {
        fmt.Printf("%v: %s", i+1, p.ShortDesc() + "\n")
	}

	invalid := func() { fmt.Println("Invalid choice. No player was picked.") }

    fmt.Print("Choose player:")
	selection, err := strconv.ParseInt(GetInput('\n'), 0, 0)
	if err != nil {
		invalid()
		return
	}
	index := int(selection - 1)
	if index < 0 || index >= len(choices) {
		invalid()
		return
	}
	picked := choices[index]
	if err := r.Pick(picked); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(picked.ShortDesc(), "was picked.", len(r.UnDrafted), "available players remaining.")
}

package command

import (
	"fmt"
	"strings"

	"github.com/davidnix/ffdraft/players"
)

func Floor(repo *players.Repo, args []string) {
	fmt.Println("FLOOR:")
	if len(args) > 0 {
		pos := strings.Join(args, "")
		PrintTable(repo.FloorByPos(pos))
	} else {
		PrintTable(repo.Floor())
	}
}

func Ceil(repo *players.Repo, args []string) {
	fmt.Println("CEILING:")
	if len(args) > 0 {
		pos := strings.Join(args, "")
		PrintTable(repo.CeilByPos(pos))
	} else {
		PrintTable(repo.Ceil())
	}
}

func Team(repo *players.Repo, args []string) {
	fmt.Println("DEPTH CHART:")
	PrintTable(repo.Team(strings.Join(args, "")))
}

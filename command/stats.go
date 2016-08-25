package command

import (
	"fmt"
	"github.com/davidnix/ffdraft/players"
	"strings"
)

func Floor(repo *players.Repo) {
	fmt.Println("FLOOR:")
	PrintTable(repo.Floor())
}

func Ceil(repo *players.Repo) {
	fmt.Println("CEILING:")
	PrintTable(repo.Ceil())
}

func Team(repo *players.Repo, args []string) {
	fmt.Println("DEPTH CHART:")
	PrintTable(repo.Team(strings.Join(args, "")))
}

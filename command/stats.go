package command

import (
	"fmt"
	"strings"

	"github.com/davidnix/ffdraft/players"
	"github.com/davidnix/ffdraft/presenter"
)

func Floor(repo *players.Repo, args []string) {
	fmt.Println("FLOOR:")
	if len(args) > 0 {
		pos := strings.Join(args, "")
		PrintTable(presenter.FloorPlayers(repo.FloorByPos(pos)))
	} else {
		PrintTable(presenter.FloorPlayers(repo.Floor()))
	}
}

func Ceil(repo *players.Repo, args []string) {
	fmt.Println("CEILING:")
	if len(args) > 0 {
		pos := strings.Join(args, "")
		PrintTable(presenter.CeilPlayers(repo.CeilByPos(pos)))
	} else {
		PrintTable(presenter.CeilPlayers(repo.Ceil()))
	}
}

func Team(repo *players.Repo, args []string) {
	fmt.Println("DEPTH CHART:")
	PrintTable(presenter.Players(repo.Team(strings.Join(args, ""))))
}

func DraftPosition(repo *players.Repo) {
	fmt.Println("Draft Position:", repo.Position)
}

package command

import (
	"log"
	"strings"

	"github.com/davidnix/ffdraft/players"
	"github.com/davidnix/ffdraft/presenter"
)

func Floor(repo *players.Repo, args []string, limit int) {
	log.Println("FLOOR:")
	if len(args) > 0 {
		pos := strings.Join(args, "")
		PrintTable(presenter.FloorPlayers(repo.FloorByPos(pos, limit)))
	} else {
		PrintTable(presenter.FloorPlayers(repo.Floor()))
	}
}

func Ceil(repo *players.Repo, args []string, limit int) {
	log.Println("CEILING:")
	if len(args) > 0 {
		pos := strings.Join(args, "")
		PrintTable(presenter.CeilPlayers(repo.CeilByPos(pos, limit)))
	} else {
		PrintTable(presenter.CeilPlayers(repo.Ceil()))
	}
}

func DepthChart(repo *players.Repo, args []string) {
	log.Println("DEPTH CHART:")
	PrintTable(presenter.Players(repo.Team(strings.Join(args, ""))))
}

func DraftPosition(repo *players.Repo) {
	log.Println("Draft Position:", repo.Position)
}

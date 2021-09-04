package command

import (
	"log"
	"strings"

	"github.com/davidnix/ffdraft/players"
	"github.com/davidnix/ffdraft/presenter"
)

func Find(repo *players.Repo, args []string) {
	found := repo.FindAll(strings.Join(args, " "))
	if len(found) == 0 {
		log.Println(errPlayerNotFound)
		return
	}
	PrintTable(presenter.Players(found))
}

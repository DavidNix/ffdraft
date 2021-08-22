package command

import (
	"strings"

	"github.com/davidnix/ffdraft/players"
	"github.com/davidnix/ffdraft/presenter"
)

func Find(repo *players.Repo, args []string) {
	found := repo.FindAll(strings.Join(args, " "))
	PrintTable(presenter.Players(found))
}

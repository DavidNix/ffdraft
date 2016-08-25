package command

import (
	"github.com/davidnix/ffdraft/players"
	"strings"
)

func Find(repo *players.Repo, args []string) {
	rows := repo.FindAll(strings.Join(args, " "))
	PrintTable(rows)
}

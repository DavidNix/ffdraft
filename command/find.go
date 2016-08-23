package command

import (
	"github.com/davidnix/ffdraft/players"
	"strings"
)

func Find(repo *players.Repo, args []string) {
	rows := repo.Find(strings.Join(args, " "))
	PrintTable(rows)
}

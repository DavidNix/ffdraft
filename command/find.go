package command

import (
    "strings"

    "github.com/davidnix/ffdraft/players"
)

func Find(repo *players.Repo, args []string) {
    rows := repo.FindAll(strings.Join(args, " "))
    PrintTable(repo, rows)
}

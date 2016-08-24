package command

import (
	"fmt"
	"github.com/davidnix/ffdraft/players"
)

func Floor(repo *players.Repo) {
	fmt.Println("FLOOR:")
	PrintTable(repo.Floor())
}

func Ceil(repo *players.Repo) {
    fmt.Println("CEILING:")
    PrintTable(repo.Ceil())
}
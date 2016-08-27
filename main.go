package main

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/davidnix/ffdraft/command"
	"github.com/davidnix/ffdraft/players"
	"log"
	"strings"
	"time"
	"flag"
)

const usage = `
--------------------------------------------------------------------------------------------------------------------
Commands:
    find, f [player name]: fuzzy finds players matching player name
    pick, p [player id]: removes player from draft pool
    unpick, u [player id]: adds player back to draft pool
    floor: print the highest floor value for available players for each position
    ceil: print the highest ceiling value for available players for each position
    help, h: print this usage text
    exit: exits this program
*By default, this program always prints the result of the floor command after every command.
--------------------------------------------------------------------------------------------------------------------
`
var api = flag.Bool("api", false, "fetch data from api")

func main() {
	fmt.Println("Welcome to fantasy football!")
	fmt.Println(usage)

    var undrafted []players.Player
    var err error

    s := startSpinner()
    flag.Parse()
    if *api {
        fmt.Println("Fetching current player data...")
        undrafted, err = players.Load()
    } else {
        undrafted, err = players.LoadFromFile(players.CacheLocation)
    }
	s.Stop()

	if err != nil {
		log.Fatal("unable to fetch player data:", err)
        return
	}

	repo := players.NewRepo(undrafted)
	fmt.Println("Loaded", len(repo.UnDrafted), "offensive players")
	command.Floor(repo, []string{})

	startInteractive(repo)

	fmt.Println("Program exited")
}

func startInteractive(repo *players.Repo) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from fatal error:", err)
            startInteractive(repo)
		}
	}()
Loop:
	for {
		input := strings.Fields(command.GetInput('\n'))
		var cmd string
		args := []string{}
		if len(input) > 0 {
			cmd = input[0]
			args = input[1:]
		}

		switch cmd {
		case "find", "f":
			command.Find(repo, args)

		case "pick", "p":
			command.Pick(repo, args)

		case "unpick", "u":
			command.UnPick(repo, args)

		case "floor", "fl":
			command.Floor(repo, args)

		case "ceil":
			command.Ceil(repo, args)

		case "team":
			command.Team(repo, args)

		case "help", "h", "usage":
			fmt.Println(usage)

		case "exit":
			break Loop

		case "":
            continue

		default:
			fmt.Println("Unrecognized command \"" + cmd + "\". Type help for usage.")
		}
	}
}

func startSpinner() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[7], 100*time.Millisecond)
	s.Start()
	return s
}

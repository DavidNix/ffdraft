package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/fatih/color"

	"github.com/briandowns/spinner"
	"github.com/davidnix/ffdraft/command"
	"github.com/davidnix/ffdraft/players"
)

const cmdUsage = `
	ffdraft -csv PATH 
`

const interactiveUsage = `
--------------------------------------------------------------------------------------------------------------------
Commands:
    find, f [name]:         fuzzy finds players matching player name
    pick, p [name]:         removes player from draft pool
    unpick, u [name]: 	    adds player back to draft pool
    keep:                   remove a player without advancing draft position (useful for keeper leagues)
    floor, fl:              print the highest floor value for available players for each position
    ceil:                   print the highest ceiling value for available players for each position
    team:                   print a team's depth chart
    position, dp:           print current draft position
    help, h:                print this interactiveUsage text
    exit, ctl+D:            exits this program
*By default, always prints the result of floor after every command.
--------------------------------------------------------------------------------------------------------------------`

// Flags
var (
	csvPath string
)

func main() {
	flag.StringVar(&csvPath, "csv", "", "PATH to csv data")
	flag.Parse()

	if e := validateFlags(); e != nil {
		fmt.Printf("Error: %v\n\nUsage:%s", e, cmdUsage)
		os.Exit(1)
	}

	fmt.Println("Welcome to fantasy football!")

	s := startSpinner()
	undrafted, err := players.LoadFromCSV(csvPath)
	if err != nil {
		log.Fatal("unable to load csv:", err)
	}
	s.Stop()

	repo := players.NewRepo(undrafted)
	color.HiGreen("Loaded %d offensive players", len(repo.UnDrafted))
	command.Floor(repo, []string{})

	fmt.Println(interactiveUsage)
	startInteractive(repo)

	fmt.Println("Program exited")
}

func validateFlags() error {
	if csvPath == "" {
		return fmt.Errorf("csv required")
	}
	return nil
}

func preventSigTerm() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	go func() {
		for _ = range ch {
			fmt.Println("Interrupt caught: ignoring. Use `exit` or ctl+D")
		}
	}()
}

func startInteractive(repo *players.Repo) {
	preventSigTerm()
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

		case "keep":
			command.Keep(repo, args)

		case "floor", "fl":
			command.Floor(repo, args)

		case "ceil":
			command.Ceil(repo, args)

		case "team":
			command.Team(repo, args)

		case "position", "dp":
			command.DraftPosition(repo)

		case "help", "h", "usage":
			fmt.Println(interactiveUsage)

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

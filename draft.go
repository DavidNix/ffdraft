package main

import (
	"log"
	"os"
	"os/signal"
	"strings"

	"github.com/davidnix/ffdraft/command"
	"github.com/davidnix/ffdraft/players"
	"github.com/fatih/color"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

var draftCmd = &cli.Command{
	Name:  "draft",
	Usage: "Start interactive draft",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "csv",
			Usage:    "(required) path to projections csv",
			Required: true,
		},
	},
	Action: interactiveAction,
}

const draftUsage = `
--------------------------------------------------------------------------------------------------------------------
Commands:
    ceil:                   print the highest ceiling value for available players for each position
    exit, ctl+D:            exits this program
    find, f [name]:         fuzzy finds players matching player name
    floor, fl:              print the highest floor value for available players for each position
    help, h:                print this help text
    keep:                   remove a player without advancing draft position (useful for keeper leagues)
    pick, p [name]:         removes player from draft pool
    position, dp:           print current draft position
    depth:                  print a team's depth chart
    unpick, u [name]: 	    adds player back to draft pool
--------------------------------------------------------------------------------------------------------------------`

func interactiveAction(ctx *cli.Context) error {
	repo, err := buildRepo(ctx.String("csv"))
	if err != nil {
		return err
	}

	log.Println("Welcome to fantasy football!")
	color.HiGreen("Loaded %d offensive players", len(repo.Available))
	command.Floor(repo, []string{})

	preventSigTerm()
	return startInteractive(repo)
}

func preventSigTerm() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	go func() {
		for range ch {
			log.Println("Interrupt caught: ignoring. Use `exit` or ctl+D")
		}
	}()
}

func startInteractive(repo *players.Repo) error {
	log.Println(draftUsage)
	for {
		in, err := command.GetInput()
		if err != nil {
			return err
		}
		input := strings.Fields(in)
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

		case "depth":
			command.DepthChart(repo, args)

		case "position", "dp":
			command.DraftPosition(repo)

		case "help", "h", "usage":
			log.Println(draftUsage)

		case "exit":
			return errors.New("user canceled")

		case "":
			continue

		default:
			log.Printf("Unrecognized command %q. Type help for usage.\n", cmd)
		}
	}
}

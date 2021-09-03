package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

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

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	app := &cli.App{
		Name:        "ffdraft",
		Usage:       "Fantasy football drafting and lineup optimization.",
		Version:     "v0.0.1",
		Description: "Relies on the good work of https://github.com/FantasyFootballAnalytics/ffanalytics",
		Commands: []*cli.Command{
			interactiveCmd,
			lineupCmd,
		},
		Authors:   []*cli.Author{{Name: "David Nix"}},
		Copyright: "MIT License",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

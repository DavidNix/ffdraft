package main

import (
	"log"

	"github.com/urfave/cli/v2"
)

var lineupCmd = &cli.Command{
	Name:        "lineup",
	Usage:       "Optimize your lineup",
	Description: "Requires using R scripts",
	Action:      lineupAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "csv",
			Usage: "path to your projections csv",
		},
	},
}

func lineupAction(ctx *cli.Context) error {
	log.Println("lineup called")
	return nil
}

package main

import (
	"github.com/davidnix/ffdraft/command"
	"github.com/urfave/cli/v2"
)

var lineupCmd = &cli.Command{
	Name:        "lineup",
	Usage:       "Optimize your lineup",
	Description: "Requires using R scripts",
	Action:      lineupAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "csv",
			Usage:    "(required) path to projections csv",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "team",
			Aliases:  []string{"t"},
			Usage:    "(required) your team name",
			Required: true,
		},
		&cli.BoolFlag{
			Name:  "floor",
			Usage: "rank by floor, if false, rank by ceil.",
			Value: true,
		},
	},
}

func lineupAction(ctx *cli.Context) error {
	repo, err := buildRepo(ctx.String("csv"))
	if err != nil {
		return err
	}
	team, err := loadTeam(ctx.String("team"))
	if err != nil {
		return err
	}
	return command.Lineup(team, repo.Available, ctx.Bool("floor"))
}

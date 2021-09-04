package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/davidnix/ffdraft/command"
	"github.com/davidnix/ffdraft/players"
	"github.com/davidnix/ffdraft/presenter"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

var teamCmd = &cli.Command{
	Name:  "team",
	Usage: "Manage a team",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "csv",
			Usage:    "(required) path to projections csv",
			Required: true,
		},
		&cli.StringFlag{
			Name:     "name",
			Aliases:  []string{"n"},
			Usage:    "(required) your team name",
			Required: true,
		},
	},
	Action: teamInteractive,
}

const teamUsage = `
--------------------------------------------------------------------------------------------------------------------
Commands:
    add [name]:             adds player to your team	
    rm [name]:              removes player from your team
    exit, ctl+D:            exits this program
    help, h:                print this help text
--------------------------------------------------------------------------------------------------------------------`

func teamInteractive(ctx *cli.Context) error {
	repo, err := buildRepo(ctx.String("csv"))
	if err != nil {
		return err
	}
	team, err := loadTeam(ctx.String("name"))
	if err != nil {
		return err
	}
	repo.SyncTeam(team)

	defer mustSaveTeam(team)

	log.Println(teamUsage)

	for {
		if len(team.Players) > 0 {
			grouped := team.Players.GroupPosition(func(p1, p2 players.Player) bool {
				return p1.Name() < p2.Name()
			}, 100)
			command.PrintTable(presenter.Team(grouped))
		}

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
		case "add":
			command.TeamAdd(repo, team, args)

		case "rm":
			command.TeamRemove(repo, team, args)

		case "exit":
			return errors.New("user canceled")

		case "help", "h", "usage":
			log.Println(draftUsage)

		case "":
			continue

		default:
			log.Printf("Unrecognized command %q. Type help for usage.\n", cmd)
		}
	}
}

func loadTeam(name string) (*players.Team, error) {
	t := &players.Team{Name: name}
	fname := t.Name + ".json"
	existing, err := os.ReadFile(fname)
	if err != nil {
		log.Printf("team file %q: %v; creating new team", fname, err)
		return t, nil
	}
	err = json.Unmarshal(existing, t)
	return t, errors.Wrap(err, fname)
}

func mustSaveTeam(t *players.Team) {
	fname := t.Name + ".json"
	log.Println("saving team to", fname)
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(fname, b, 0660)
	if err != nil {
		panic(err)
	}
}

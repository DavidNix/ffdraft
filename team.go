package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"github.com/davidnix/ffdraft/command"
	"github.com/davidnix/ffdraft/players"
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
	ceil:                   print the highest ceiling value for available players for each position
	depth, team:            print a team's depth chart
	exit, ctl+D:            exits this program
	find, f [name]:         fuzzy finds players matching player name
	floor, fl:              sort by highest floor value for your team
	help, h:                print this help text
	rm, drop [name]:        removes player from your team
	show, s:                show your team grouped by position
	waiver, w:              show available sorted by highest floor value
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
	log.Println("Loaded", len(repo.Available), "players")
	repo.Sync(team.Players)
	log.Println(len(repo.Claimed), "players on your team")
	log.Println(len(repo.Available), "players remaining")

	defer mustSaveTeam(team, repo)

	log.Println(teamUsage)

	command.Lineup(repo)

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
		case "add":
			command.Pick(repo, args)
			command.Lineup(repo)

		case "rm", "drop":
			command.UnPick(repo, args)
			command.Lineup(repo)

		case "exit":
			return errors.New("user canceled")

		case "floor", "fl":
			command.LineupFloor(repo)

		case "ceil":
			command.LineupCeil(repo)

		case "depth", "team":
			command.Team(repo, args)

		case "find", "f":
			command.Find(repo, args)

		case "help", "h", "usage":
			log.Println(teamUsage)

		case "waiver", "w":
			command.Floor(repo, args)

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

func mustSaveTeam(t *players.Team, repo *players.Repo) {
	t.Sync(repo.Claimed)
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

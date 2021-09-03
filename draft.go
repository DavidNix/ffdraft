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
)

func interactiveDraft(cfg cmdConfig) error {
	if cfg.csvPath == "" {
		return errors.New("csv path required")
	}

	f, err := os.Open(cfg.csvPath)
	if err != nil {
		return err
	}
	defer f.Close()

	log.Println("Welcome to fantasy football!")

	undrafted, err := players.LoadFromCSV(f)
	if err != nil {
		return err
	}

	repo := players.NewRepo(undrafted)
	color.HiGreen("Loaded %d offensive players", len(repo.UnDrafted))
	command.Floor(repo, []string{})

	log.Println(interactiveUsage)
	startInteractive(repo)

	log.Println("program exited")
	return nil
}

func preventSigTerm() {
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	go func() {
		for range ch {
			log.Println("Interrupt caught: ignoring. Use `exit` or ctl+D")
		}
	}()
}

func startInteractive(repo *players.Repo) {
	preventSigTerm()
	defer func() {
		if err := recover(); err != nil {
			log.Println("Recovered from fatal error:", err)
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
			log.Println(interactiveUsage)

		case "exit":
			break Loop

		case "":
			continue

		default:
			log.Println("Unrecognized command \"" + cmd + "\". Type help for usage.")
		}
	}
}

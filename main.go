package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	app := &cli.App{
		Name:        "ffdraft",
		Usage:       "Fantasy football drafting and lineup optimization.",
		Version:     "v0.0.1",
		Description: "Relies on the good work of https://github.com/FantasyFootballAnalytics/ffanalytics",
		Commands: []*cli.Command{
			draftCmd,
			teamCmd,
		},
		Authors:   []*cli.Author{{Name: "David Nix"}},
		Copyright: "MIT License",
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

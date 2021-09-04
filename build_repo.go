package main

import (
	"os"

	"github.com/davidnix/ffdraft/players"
	"github.com/pkg/errors"
)

func buildRepo(csvPath string) (*players.Repo, error) {
	if csvPath == "" {
		return nil, errors.New("csv path required")
	}

	f, err := os.Open(csvPath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	available, err := players.LoadFromCSV(f)
	if err != nil {
		return nil, err
	}

	return players.NewRepo(available), nil
}

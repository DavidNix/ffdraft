package main

import (
	"os"

	"github.com/davidnix/ffdraft/players"
	"github.com/pkg/errors"
)

func buildRepo(path string) (*players.Repo, error) {
	if path == "" {
		return nil, errors.New("csv path required")
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	undrafted, err := players.LoadFromCSV(f)
	if err != nil {
		return nil, err
	}

	return players.NewRepo(undrafted), nil
}

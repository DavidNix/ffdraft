package main

import (
	"fmt"
	"github.com/briandowns/spinner"
	"github.com/davidnix/ffdraft/datasource"
	"log"
	"time"
)

func main() {
	fmt.Println("Fetching current player data...")

	s := spinner.New(spinner.CharSets[7], 100*time.Millisecond)
	s.Start()
	defer s.Stop()
	players, err := datasource.LoadPlayers()
	if err != nil {
		log.Fatal("unable to fetch player data: ", err)
	}

	fmt.Println("total players", len(players))

	fmt.Println("Program exited")
}

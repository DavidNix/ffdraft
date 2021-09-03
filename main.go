package main

import (
	"flag"
	"log"
	"os"
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

type cmdConfig struct {
	csvPath string
}

func main() {
	log.SetFlags(0)
	log.SetOutput(os.Stdout)

	var cfg cmdConfig

	flag.StringVar(&cfg.csvPath, "csv", "", "PATH to csv data")
	flag.Parse()

	if err := interactiveDraft(cfg); err != nil {
		log.Fatalln(err)
	}
}

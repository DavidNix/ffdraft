package main

import (
	"fmt"
	"github.com/briandowns/spinner"
	//"github.com/davidnix/ffdraft/datasource"
	"bufio"
	"log"
	"os"
	"strings"
	"time"
)

const usage = `
Commands:
    find, f [player name]: fuzzy finds players matching player name
    pick, p [player name]: fuzzy finds players and asks which player to pick, removing them from the draft pool
    unpick, u [player name]: fuzzy finds players and asks which player to unpick, adding them back to the draft pool
    floor: print the highest floor value for available players for each position
    ceil: print the highest ceiling value for available players for each position
    help, h: print this usage text
    exit: exits this program
*By default, this program always prints the result of the floor command after every command.
`

func main() {
	fmt.Println("Welcome to fantasy football!")
	fmt.Println(usage, "\n")
	fmt.Println("Fetching current player data...")

	s := startSpinner()
	//players, err := datasource.LoadPlayers()
	s.Stop()
	//if err != nil {
	//	log.Fatal("unable to fetch player data:", err)
	//}
	//fmt.Println("total players", len(players))
Loop:
	for {
		input := strings.Fields(getInput())
		var cmd string
		var args []string
		if len(input) > 0 {
			cmd = input[0]
			args = input[1:]
			fmt.Println("args", args)
		}

		switch cmd {
		case "find", "f":
			fmt.Println("find a player")

		case "pick", "p":
			fmt.Println("pick from a list of players")

		case "unpick", "u":
			fmt.Println("unpick from a list of players")

		case "floor":
			fmt.Println("find the top floor players")

		case "ceil":
			fmt.Println("find the top ceiling players")

		case "help", "h", "usage":
			fmt.Println(usage, "\n")

		case "exit":
			break Loop

		case "":
			break

		default:
			fmt.Println("Unrecognized command \"" + cmd + "\". Type help for usage.")
		}
	}

	fmt.Println("Program exited")
}

func getInput() string {
	fmt.Print("Command: ")
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		log.Fatal("unable to read input:", err)
	}
	return strings.TrimSpace(line)
}

func startSpinner() *spinner.Spinner {
	s := spinner.New(spinner.CharSets[7], 100*time.Millisecond)
	s.Start()
	return s
}

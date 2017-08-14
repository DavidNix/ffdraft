package command

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func GetInput(delim byte) string {
	fmt.Print("> ")
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString(delim)
	if err != nil {
		log.Fatal("unable to read input:", err)
	}
	return strings.TrimSpace(line)
}

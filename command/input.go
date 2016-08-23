package command

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func GetInput() string {
	fmt.Print("> ")
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	if err != nil {
		log.Fatal("unable to read input:", err)
	}
	return strings.TrimSpace(line)
}

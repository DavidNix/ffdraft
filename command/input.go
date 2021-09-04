package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput() (string, error) {
	fmt.Print("> ")
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString('\n')
	return strings.TrimSpace(line), err
}

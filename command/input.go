package command

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInput(delim byte) (string, error) {
	fmt.Print("> ")
	in := bufio.NewReader(os.Stdin)
	line, err := in.ReadString(delim)
	return strings.TrimSpace(line), err
}

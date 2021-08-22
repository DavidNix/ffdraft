package presenter

import (
	"fmt"
	"strconv"

	"github.com/fatih/color"
)

type Float float64

func (f Float) String() string { return fmt.Sprintf("%.1f", f) }

type Int int

func (i Int) String() string { return strconv.Itoa(int(i)) }

type Injury string

func (injury Injury) String() string {
	i := string(injury)
	switch i {
	case "":
		return i
	case "Q":
		return color.YellowString(i)
	default:
		return color.RedString(i)
	}
}

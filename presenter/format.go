package presenter

import (
	"fmt"
	"strconv"
)

type Float float64

func (f Float) String() string { return fmt.Sprintf("%.1f", f) }

type Int int

func (i Int) String() string { return strconv.Itoa(int(i)) }

package players

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBy_Sort(t *testing.T) {
	plys := []Player{
		{Name: "Peyton Manning", Floor: 1},
		{Name: "Russel Wilson", Floor: 2},
		{Name: "Emmit Smith", Floor: 4},
		{Name: "Khalil Mack", Floor: 10},
		{Name: "David Carr", Floor: 3},
	}

	floor := func(p1, p2 Player) bool {
		return p1.Floor > p2.Floor
	}
	By(floor).Sort(plys)

	assert.Equal(t, plys[0].Name, "Khalil Mack")
	assert.Equal(t, plys[1].Name, "Emmit Smith")
	assert.Equal(t, plys[2].Name, "David Carr")
	assert.Equal(t, plys[3].Name, "Russel Wilson")
	assert.Equal(t, plys[4].Name, "Peyton Manning")
}

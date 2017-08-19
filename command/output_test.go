package command

import (
	"testing"

	"github.com/davidnix/ffdraft/players"
	"github.com/stretchr/testify/assert"
)

func TestPrintTable(t *testing.T) {
	assert.NotPanics(t, func() {
		rows := Rows{players.Player{}}
		PrintTable(rows)
	})
}

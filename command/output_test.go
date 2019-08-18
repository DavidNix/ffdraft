package command

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/davidnix/ffdraft/players"
)

func TestPrintTable(t *testing.T) {
	require.NotPanics(t, func() {
		rows := Rows{players.Player{}}
		PrintTable(rows)
	})
}

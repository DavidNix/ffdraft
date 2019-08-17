package players

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayer_Row(t *testing.T) {
	row := Player{ID: 1}.Row(0)
	require.Len(t, row, 18)
}

func TestPlayer_Row_blank(t *testing.T) {
	row := Player{}.Row(0)
	require.Len(t, row, 18)
}

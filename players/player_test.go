package players

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayer_String(t *testing.T) {
	p := Player{NameFirst: "Bob", NameLast: "Smith", Position: "DST", Team: "DAL"}

	require.Equal(t, "Bob Smith DST DAL", p.String())
}

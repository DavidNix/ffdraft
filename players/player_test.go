package players

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPlayer_String(t *testing.T) {
	p := Player{Name: "Bob", Position: "DST", Team: "DAL"}

	require.Equal(t, "Bob DST DAL", p.String())
}

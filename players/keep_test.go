package players

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepo_Keep(t *testing.T) {
	p := Player{ID: 1}
	r := &Repo{
		UnDrafted: []Player{p},
	}
	var err error

	p.ID = 2
	err = r.Keep(p)
	require.Error(t, err)

	require.Equal(t, 0, len(r.Drafted))
	require.Equal(t, 1, len(r.UnDrafted))
	require.Equal(t, 0, r.Position)

	p.ID = 1
	err = r.Keep(p)
	require.NoError(t, err)

	require.Equal(t, 0, len(r.UnDrafted))
	require.Equal(t, 1, len(r.Drafted))
	require.Equal(t, 0, r.Position)
}

package players

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepo_Pick(t *testing.T) {
	p := Player{Name: "1"}
	r := &Repo{
		UnDrafted: []Player{p},
	}
	var err error

	p.Name = "2"
	err = r.Pick(p)
	require.Error(t, err)

	require.Equal(t, 0, len(r.Drafted))
	require.Equal(t, 1, len(r.UnDrafted))
	require.Equal(t, 0, r.Position)

	p.Name = "1"
	err = r.Pick(p)
	require.NoError(t, err)
	require.Equal(t, 0, len(r.UnDrafted))
	require.Equal(t, 1, len(r.Drafted))
	require.Equal(t, 1, r.Position)
}

func TestRepo_UnPick(t *testing.T) {
	p := Player{Name: "1"}
	r := &Repo{
		Position: 3,
		Drafted:  []Player{p},
	}

	var err error

	p.Name = "2"
	err = r.UnPick(p)

	require.Error(t, err)
	require.Equal(t, 1, len(r.Drafted))
	require.Equal(t, 0, len(r.UnDrafted))
	require.Equal(t, 3, r.Position)

	p.Name = "1"
	err = r.UnPick(p)

	require.NoError(t, err)
	require.Equal(t, 0, len(r.Drafted))
	require.Equal(t, 1, len(r.UnDrafted))
	require.Equal(t, 2, r.Position)
}

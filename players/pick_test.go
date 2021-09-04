package players

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepo_Pick(t *testing.T) {
	p := Player{ID: 1}
	r := &Repo{
		Available: []Player{{ID: 1, ADP: 45}},
	}
	var err error

	p.ID = 2
	err = r.Pick(p)
	require.Error(t, err)

	require.Equal(t, 0, len(r.Claimed))
	require.Equal(t, 1, len(r.Available))
	require.Equal(t, 0, r.Position)

	p.ID = 1
	err = r.Pick(p)
	require.NoError(t, err)
	require.Equal(t, 0, len(r.Available))
	require.Equal(t, 1, len(r.Claimed))
	require.Equal(t, 1, r.Position)
}

func TestRepo_UnPick(t *testing.T) {
	p := Player{ID: 1}
	r := &Repo{
		Position: 3,
		Claimed:  []Player{{ID: 1, ADP: 45}},
	}

	var err error

	p.ID = 2
	err = r.UnPick(p)

	require.Error(t, err)
	require.Equal(t, 1, len(r.Claimed))
	require.Equal(t, 0, len(r.Available))
	require.Equal(t, 3, r.Position)

	p.ID = 1
	err = r.UnPick(p)

	require.NoError(t, err)
	require.Equal(t, 0, len(r.Claimed))
	require.Equal(t, 1, len(r.Available))
	require.Equal(t, 2, r.Position)
}

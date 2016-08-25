package players

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepo_Pick(t *testing.T) {
	p := Player{ID: 10}
	r := &Repo{
		UnDrafted: []Player{p},
	}
	var err error

	p.ID = 11
	err = r.Pick(p)

	assert.Error(t, err)
	assert.Equal(t, len(r.Drafted), 0)
	assert.Equal(t, len(r.UnDrafted), 1)

	p.ID = 10
	err = r.Pick(p)

	assert.Nil(t, err)
	assert.Equal(t, len(r.UnDrafted), 0)
	assert.Equal(t, len(r.Drafted), 1)
}

func TestRepo_UnPick(t *testing.T) {
	p := Player{ID: 1}
	r := &Repo{
		Drafted: []Player{p},
	}

	var err error

	p.ID = 2
	err = r.UnPick(p)

	assert.Error(t, err)
	assert.Equal(t, len(r.Drafted), 1)
	assert.Equal(t, len(r.UnDrafted), 0)

	p.ID = 1
	err = r.UnPick(p)

	assert.Nil(t, err)
	assert.Equal(t, len(r.Drafted), 0)
	assert.Equal(t, len(r.UnDrafted), 1)
}

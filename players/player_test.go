package players

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayer_Row(t *testing.T) {
	row := Player{ID: 1}.Row()
	assert.Len(t, row, 18)
}

func TestPlayer_Row_blank(t *testing.T) {
	row := Player{}.Row()
	assert.Len(t, row, 18)
}

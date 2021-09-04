package presenter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTeam_Headers(t *testing.T) {
	pres := Players{}
	h := pres.Headers()
	require.NotEmpty(t, h)
}

func TestTeam_Rows(t *testing.T) {
	pres := make(Players, 2)
	require.Len(t, pres.Rows(), 2)
	require.Equal(t, len(pres.Headers()), len(pres.Rows()[0]))
}

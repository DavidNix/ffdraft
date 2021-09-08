package presenter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPresenters(t *testing.T) {
	const count = 2
	type iface interface {
		Headers() []string
		Rows() [][]string
	}
	for _, tt := range []struct {
		Presenter iface
	}{
		{make(Team, count)},
		{make(Players, count)},
		{make(CeilPlayers, count)},
		{make(FloorPlayers, count)},
	} {
		h := tt.Presenter.Headers()
		msg := fmt.Sprintf("%+v", tt)

		require.NotEmpty(t, h, msg)
		require.Len(t, tt.Presenter.Rows(), 2, msg)
		require.Equal(t, len(tt.Presenter.Headers()), len(tt.Presenter.Rows()[0]), msg)
	}
}

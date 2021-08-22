package command

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type mockPresenter struct{}

func (m mockPresenter) Headers() []string { return nil }
func (m mockPresenter) Rows() [][]string  { return nil }

func TestPrintTable(t *testing.T) {
	require.NotPanics(t, func() {
		PrintTable(mockPresenter{})
	})
}

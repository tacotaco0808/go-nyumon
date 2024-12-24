package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStation02(t *testing.T) {
	t.Run("Swap関数が正しく実装できている", func(t *testing.T) {
		x, y := 10, 100

		Swap(&x, &y)

		require.Equal(t, 100, x)
		require.Equal(t, 10, y)
	})
}

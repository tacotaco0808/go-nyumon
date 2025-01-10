package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStation04(t *testing.T) {
	t.Parallel()

	t.Run("正しく四角形の面積を求められる", func(t *testing.T) {
		t.Parallel()

		var width, height float64 = 10.5, 34
		rectangle := Rectangle{
			Width:  width,
			Height: height,
		}

		result := rectangle.Area()
		expected := width * height

		require.Equal(t, expected, result)
	})

	t.Run("正しく円の面積を求められる", func(t *testing.T) {
		t.Parallel()

		radius := 22.2
		circle := Circle{
			Radius: radius,
		}

		result := circle.Area()
		expected := radius * radius * math.Pi

		require.Equal(t, expected, result)
	})
}

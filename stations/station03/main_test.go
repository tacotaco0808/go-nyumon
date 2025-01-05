package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStation03(t *testing.T) {
	t.Run("構造体Bookを正しく定義できている", func(t *testing.T) {
		book := Book{
			Title:      "タイトル",
			Author:     "著者",
			Price:      0,
			Categories: []string{"カテゴリ1", "カテゴリ2"},
		}

		require.NotNil(t, book)
	})

	t.Run("Total関数を正しく実装できている", func(t *testing.T) {
		books := []Book{
			{
				Price: 100,
			},
			{
				Price: 200,
			},
			{
				Price: 300,
			},
		}

		totalPrice := Total(books)
		expected := 600

		require.Equal(t, expected, totalPrice)
	})

	t.Run("ToMap関数を正しく実装できている", func(t *testing.T) {
		books := []Book{
			{
				Title:  "A",
				Author: "田中 太郎",
			},
			{
				Title:  "B",
				Author: "佐藤 花子",
			},
			{
				Title:  "C",
				Author: "鈴木 次郎",
			},
		}

		authorBooksMap := ToMap(books)

		require.Equal(t, "A", authorBooksMap["田中 太郎"].Title)
		require.Equal(t, "B", authorBooksMap["佐藤 花子"].Title)
		require.Equal(t, "C", authorBooksMap["鈴木 次郎"].Title)
	})
}

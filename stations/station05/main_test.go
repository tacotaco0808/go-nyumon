package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStation05(t *testing.T) {

	tests := []struct {
		Name     string
		FilePath string
		Expected string
	}{
		{
			Name:     "ファイルが存在する場合に正しいメッセージが出力される",
			FilePath: "sample.txt",
			Expected: "ファイルが見つかりました",
		},
		{
			Name:     "ファイルが存在しない場合に正しいメッセージが表示される",
			FilePath: "xxx",
			Expected: "ファイルが見つかりませんでした",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			var buf bytes.Buffer
			r, w, _ := os.Pipe()
			os.Stdout = w

			CheckFileExist(test.FilePath)

			err := w.Close()
			if err != nil {
				return
			}
			_, err = buf.ReadFrom(r)
			if err != nil {
				return
			}

			got := buf.String()

			require.Equal(t, test.Expected, got)
		})
	}

}

package tests

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestStation06(t *testing.T) {
	t.Run("`go run ./app/main.go` を実行できる", func(t *testing.T) {
		rootDir, _ := os.Getwd()
		parentDir := filepath.Dir(rootDir)
		cmd := exec.Command("go", "run", filepath.Join(parentDir, "app", "main.go"))
		cmd.Dir = parentDir
		output, err := cmd.CombinedOutput()

		if err != nil {
			t.Errorf("コマンドの実行に失敗しました: %v\n出力: %s", err, output)
		}

		if len(output) == 0 {
			t.Error("コマンドの出力が空です")
		}
	})

	t.Run("cobra パッケージが追加されている", func(t *testing.T) {
		rootDir, _ := os.Getwd()
		parentDir := filepath.Dir(rootDir)
		cmd := exec.Command("go", "list", "-m", "github.com/spf13/cobra")
		cmd.Dir = parentDir

		_, err := cmd.CombinedOutput()
		if err != nil {
			t.Error("github.com/spf13/cobraパッケージが見つかりません")
		}
	})
}

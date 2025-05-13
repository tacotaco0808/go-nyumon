package tests

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestStation07(t *testing.T) {
	t.Run("`go run ./app/main.go random` を実行できる", func(t *testing.T) {
		rootDir, _ := os.Getwd()
		parentDir := filepath.Dir(rootDir)
		cmd := exec.Command("go", "run", filepath.Join(parentDir, "app", "main.go"), "random")
		cmd.Dir = parentDir
		output, err := cmd.CombinedOutput()

		if err != nil {
			t.Errorf("コマンドの実行に失敗しました: %v\n出力: %s", err, output)
		}

		outputStr := strings.TrimSpace(string(output))

		if !strings.HasPrefix(outputStr, "https://images.dog.ceo/breeds/") {
			t.Errorf("不正なURL形式です: %s", outputStr)
		}

		if !(strings.HasSuffix(outputStr, ".jpg") ||
			strings.HasSuffix(outputStr, ".jpeg") ||
			strings.HasSuffix(outputStr, ".png")) {
			t.Errorf("不正な画像ファイル形式です: %s", outputStr)
		}
	})
}

package tests

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestStation08(t *testing.T) {
	t.Parallel()

	t.Run("`go run ./app/main.go random` を実行できる", func(t *testing.T) {
		t.Parallel()

		rootDir, _ := os.Getwd()
		parentDir := filepath.Dir(rootDir)
		cmd := exec.Command("go", "run", filepath.Join(parentDir, "app", "main.go"), "random")
		cmd.Dir = parentDir
		output, err := cmd.CombinedOutput()

		if err != nil {
			t.Errorf("コマンドの実行に失敗しました: %v\n出力: %s", err, output)
		}

		urls := strings.Split(strings.TrimSpace(string(output)), "\n")
		if len(urls) != 1 {
			t.Errorf("デフォルトで1枚の画像URLが期待されますが、%d枚返されました", len(urls))
		}

		validateURL(t, urls[0])
	})

	t.Run("`go run ./app/main.go random --images 3` を実行できる", func(t *testing.T) {
		t.Parallel()

		rootDir, _ := os.Getwd()
		parentDir := filepath.Dir(rootDir)
		cmd := exec.Command("go", "run", filepath.Join(parentDir, "app", "main.go"), "random", "--images", "3")
		cmd.Dir = parentDir
		output, err := cmd.CombinedOutput()

		if err != nil {
			t.Errorf("コマンドの実行に失敗しました: %v\n出力: %s", err, output)
		}

		urls := strings.Split(strings.TrimSpace(string(output)), "\n")
		if len(urls) != 3 {
			t.Errorf("3枚の画像URLが期待されますが、%d枚返されました", len(urls))
		}

		for _, url := range urls {
			validateURL(t, url)
		}
	})

	t.Run("`go run ./app/main.go dog-cli random -i 2` を実行できる", func(t *testing.T) {
		t.Parallel()

		rootDir, _ := os.Getwd()
		parentDir := filepath.Dir(rootDir)
		cmd := exec.Command("go", "run", filepath.Join(parentDir, "app", "main.go"), "random", "-i", "2")
		cmd.Dir = parentDir
		output, err := cmd.CombinedOutput()

		if err != nil {
			t.Errorf("コマンドの実行に失敗しました: %v\n出力: %s", err, output)
		}

		urls := strings.Split(strings.TrimSpace(string(output)), "\n")
		if len(urls) != 2 {
			t.Errorf("2枚の画像URLが期待されますが、%d枚返されました", len(urls))
		}

		for _, url := range urls {
			validateURL(t, url)
		}
	})
}

func validateURL(t *testing.T, url string) {
	t.Helper()

	if !strings.HasPrefix(url, "https://images.dog.ceo/breeds/") {
		t.Errorf("不正なURL形式です: %s", url)
	}

	if !(strings.HasSuffix(url, ".jpg") ||
		strings.HasSuffix(url, ".jpeg") ||
		strings.HasSuffix(url, ".png")) {
		t.Errorf("不正な画像ファイル形式です: %s", url)
	}
}

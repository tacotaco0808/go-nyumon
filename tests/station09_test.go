package tests

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestStation09(t *testing.T) {
	t.Parallel()

	t.Run("犬種を指定して画像を取得できる", func(t *testing.T) {
		t.Parallel()

		rootDir, _ := os.Getwd()
		parentDir := filepath.Dir(rootDir)
		cmd := exec.Command("go", "run", filepath.Join(parentDir, "app", "main.go"), "breed", "husky")
		cmd.Dir = parentDir
		output, err := cmd.CombinedOutput()

		if err != nil {
			t.Errorf("コマンドの実行に失敗しました: %v\n出力: %s", err, output)
		}

		outputStr := strings.TrimSpace(string(output))
		if !strings.HasPrefix(outputStr, "https://images.dog.ceo/breeds/husky/") {
			t.Errorf("不正なURL形式です: %s", outputStr)
		}
	})

	t.Run("指定されたエラーメッセージを出力できる", func(t *testing.T) {
		t.Parallel()

		rootDir, _ := os.Getwd()
		parentDir := filepath.Dir(rootDir)
		cmd := exec.Command("go", "run", filepath.Join(parentDir, "app", "main.go"), "breed")
		cmd.Dir = parentDir
		output, err := cmd.CombinedOutput()

		if err == nil {
			t.Error("エラーが発生すべきですが、成功しました")
		}

		outputStr := strings.TrimSpace(string(output))
		expectedError := "Error: 犬種が指定されていません"
		if !strings.Contains(outputStr, expectedError) {
			t.Errorf("期待されるエラーメッセージ「%s」が出力されませんでした。実際の出力: %s",
				expectedError, outputStr)
		}
	})
}

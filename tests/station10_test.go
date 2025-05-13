package tests

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

func TestStation10(t *testing.T) {
	t.Parallel()

	t.Run("`go run ./app/main.go breed-list` を実行できる", func(t *testing.T) {
		t.Parallel()

		rootDir, _ := os.Getwd()
		parentDir := filepath.Dir(rootDir)
		cmd := exec.Command("go", "run", filepath.Join(parentDir, "app", "main.go"), "breed-list")
		cmd.Dir = parentDir
		output, err := cmd.CombinedOutput()

		if err != nil {
			t.Errorf("コマンドの実行に失敗しました: %v\n出力: %s", err, output)
		}

		breeds := strings.Split(strings.TrimSpace(string(output)), "\n")

		if len(breeds) == 0 {
			t.Error("犬種一覧が空です")
		}

		expectedBreeds := []string{"husky", "labrador", "poodle"}
		for _, expectedBreed := range expectedBreeds {
			found := false
			for _, breed := range breeds {
				if strings.Contains(strings.ToLower(breed), expectedBreed) {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("期待される犬種 '%s' が一覧に含まれていません", expectedBreed)
			}
		}

		breedMap := make(map[string]bool)
		for _, breed := range breeds {
			if breedMap[breed] {
				t.Errorf("犬種 '%s' が重複しています", breed)
			}
			breedMap[breed] = true
		}
	})
}

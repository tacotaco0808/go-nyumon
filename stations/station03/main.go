package main

import "fmt"

func main() {
	books := []Book{} // 値は自由に定義してください

	totalPrice := Total(books)
	fmt.Println("合計", totalPrice, "円")

	authorBooksMap := ToMap(books)
	authorName := ""
	fmt.Println("著者: ", authorName, ", タイトル: ", authorBooksMap[authorName].Title)
}

// 構造体 Book を定義 ---

// フィールド
// Title ... 文字列、本のタイトル
// Author ... 文字列、著者
// Price ... 整数、価格
// Categories ... 文字列のスライス、カテゴリ

// --------------------

// Bookの配列を引数とし、本の合計価格を戻り値とする関数 `Total` を実装する
func Total() {}

// 3. キーを「著者名 (Author)」、値を構造体 Book とするマップを戻り値とする関数 `ToMap` を実装する
func ToMap() {}

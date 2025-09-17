package main

import "fmt"
type Book struct{
	Title string
	Author string
	Price int 
	Categories []string
}
func main() {
	books := []Book{
		{
			Title:      "タイトル",
			Author:     "著者",
			Price:      30,
			Categories: []string{"カテゴリ1", "カテゴリ2"},
		},
		{
			Title:      "hoge",
			Author:     "asdf",
			Price:      220,
			Categories: []string{"カテゴリ4", "カテゴリ5"},
		},
	} // 値は自由に定義してください

	totalPrice := Total(books)
	fmt.Println("合計", totalPrice, "円")

	authorBooksMap := ToMap(books)
	authorName := "著者"
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
func Total(books []Book) int{
	sum := 0
	for _,book:=range books{
		sum += book.Price
	}
	return sum
}

// 3. キーを「著者名 (Author)」、値を構造体 Book とするマップを戻り値とする関数 `ToMap` を実装する
func ToMap(books []Book) map[string]Book{
	myMap := make(map[string]Book)
	for _,book:=range books{
		myMap[book.Author] = book
	}
	return myMap
}

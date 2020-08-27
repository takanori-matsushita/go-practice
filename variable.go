package main

import (
	"fmt"
)

func main() {
	// 変数の定義
	var num int
	num = 1
	fmt.Println(num)
	two()
	three()
}

func two() {
	// データ型は省略できる
	var num = 2
	fmt.Println(num)
}

func three() {
	// varも:=で省略できる
	num := 3
	fmt.Println(num)
}

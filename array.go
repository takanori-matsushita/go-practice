package main

import (
	"fmt"
)

func main() {
	// 配列の数とデータ型を指定する
	array := [3]string{"lawson", "family mart", "seven eleven"}
	// 配列の数を省略する
	a := [...]string{"かねひで", "サンエー", "ユニオン"}
	fmt.Println(array[0])
	fmt.Println(array[1])
	fmt.Println(array[2])

	fmt.Println(a[0])
	fmt.Println(a[1])
	fmt.Println(a[2])
}

package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("end")
	for i := 0; i <= 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("start")
}

package main

import (
	"fmt"
)

func print1() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

func print2() {
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}
}

func print3() {
	i := 0
	for {
		fmt.Println(i)
		if i == 10 {
			break
		}
		i++
	}
}

func main() {
	fmt.Println("print1")
	print1()
	fmt.Println("print2")
	print2()
	fmt.Println("print3")
	print3()
}

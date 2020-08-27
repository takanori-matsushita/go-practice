package main

import (
	"fmt"
)

func cal(x, y int) (sum int) {
	sum = x + y
	return sum
}

func main() {
	result := cal(10, 5)
	fmt.Println(result)
}

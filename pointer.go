package main

import (
	"fmt"
)

func main() {
	name := "taka"
	var point *string = &name
	fmt.Println(name)
	fmt.Println(&name)
	fmt.Println(point)
	*point = "yuki"
	fmt.Println(*point)
	fmt.Println(point)
}

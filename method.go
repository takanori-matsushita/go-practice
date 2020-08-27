package main

import (
	"fmt"
)

type User struct {
	name string
}

func (u User) cal(weight, height float64) (result float64) {
	result = weight / height / height * 10000
	return
}

func main() {
	u := User{"taka"}
	fmt.Println(u.name, u.cal(65, 174))
}

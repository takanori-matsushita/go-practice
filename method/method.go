package main

import "fmt"

type User struct {
	name string
}

func (s User) cal(weight, height float64) (result float64) {
	result = weight / height / height * 10000
	return
}

func main() {
	user := User{"taka"}

	fmt.Println(user.name, user.cal(68, 174))
}

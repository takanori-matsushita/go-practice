package main

import "fmt"

type User struct {
	gender string
	age    int
}

func main() {
	var u User
	u.gender = "male"
	u.age = 20
	fmt.Println(u)

	U := User{gender: "female", age: 30}
	fmt.Println(U)
}

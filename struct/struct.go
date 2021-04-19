package main

import "fmt"

type User struct {
	gender string
	age    int
}

func main() {
	var m User
	m.gender = "male"
	m.age = 20
	fmt.Println(m)
	f := User{"female", 30}
	fmt.Println(f)
}

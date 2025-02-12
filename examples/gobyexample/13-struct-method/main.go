package main

import "fmt"

type user struct {
	name string
	age  int
}

func (u *user) getUserName() string {
	return u.name
}

func (u *user) getUserAge() int {
	return u.age
}

func (u *user) setUserAge(age int) {
	u.age = age
}

func main() {
	u := user{"John", 15}

	fmt.Println(u.getUserName())

	fmt.Println(u.getUserAge())

	u.setUserAge(30)

	fmt.Println(u.getUserAge())
}

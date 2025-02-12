package main

import "fmt"

type user struct {
	name     string
	password string
}

func main() {

	a := user{name: "wang", password: "123456"}
	b := user{"wang", "123456"}
	c := user{name: "wang"}
	fmt.Println(a, b, c)

	var d user
	d.name = "biz"
	d.password = "123456"

	fmt.Println(d)

	fmt.Println(checkPassword(a, "123456"))
	fmt.Println(checkPassword2(&a, "123456"))

}

func checkPassword(u user, password string) bool {
	return u.password == password
}

func checkPassword2(u *user, password string) bool {
	return u.password == password
}

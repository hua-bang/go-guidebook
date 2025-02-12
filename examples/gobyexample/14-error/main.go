package main

import (
	"errors"
	"fmt"
)

type user struct {
	name     string
	password string
}

func findUser(users []user, name string) (v *user, err error) {
	for _, u := range users {
		if u.name == name {
			return &u, nil
		}
	}

	return nil, errors.New("user not found")
}

func main() {
	u, err := findUser([]user{{"wang", "1024"}}, "wang")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.name)

	if u, err := findUser([]user{{"wang", "1024"}}, "123"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u.name)
	}
}

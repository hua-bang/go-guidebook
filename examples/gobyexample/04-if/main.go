package main

import "fmt"

func main() {
	if 7 % 2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if i := 3; i == 3 {
		fmt.Println("i is 3")
	}

	c := 4;

	if c == 4 && 1 == 2 {
		fmt.Println("c == 4 && 1 == 2, true")
	} else {
		fmt.Println("c == 4 && 1 == 2, false")
	}

	if c == 4 || 1 == 2 {
		fmt.Println("c == 4 || 1 == 2, true")
	} else {
		fmt.Println("c == 4 || 1 == 2, false")
	}
}
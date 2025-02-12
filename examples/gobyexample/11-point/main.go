package main

import "fmt"

func add(n int) {
	n += 2
}

func addPtr(n *int) {
	*n += 2
}

func main() {
	n := 5
	add(n)
	fmt.Println(n)

	addPtr(&n)
	fmt.Println(n)
}

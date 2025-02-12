package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)
	fmt.Println(len(m))

	fmt.Println(m["one"])
	fmt.Println(m["two"])
	fmt.Println(m["unknown"])

	r, ok := m["unknown"]
	fmt.Println(r, ok)
	r, ok = m["two"]
	fmt.Println(r, ok)

	delete(m, "one")

	m2 := map[string]int{"one": 1, "two": 2}
	var m3 = map[string]int{"one": 1, "two": 2}
	fmt.Println(m2)
	fmt.Println(m3)
}

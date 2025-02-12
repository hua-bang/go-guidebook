package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func exists(m map[string]string, key string) (v string, ok bool) {
	v, ok = m[key]
	return v, ok
}

func main() {
	res := add(1, 2)
	fmt.Println("res", res)

	v, ok := exists(map[string]string{"a": "A"}, "a")
	fmt.Println(v, ok)
	v, ok = exists(map[string]string{"a": "A"}, "b")
	fmt.Println(v, ok)
}

package main

import "fmt"

// 数组相关的
func main() {
	nums := []int{2, 3, 4}
	sum := 0

	for i, num := range nums {
		sum += num
		if num == 2 {
			fmt.Println("index:", i, "num:", num) // index: 0 num: 2
		}
	}
	fmt.Println("sum:", sum)

	m := map[string]string{"a": "A", "b": "B"}
	for k, v := range m {
		fmt.Println("key:", k, "value:", v)
	}

	for k := range m {
		fmt.Println("key:", k)
	}
}
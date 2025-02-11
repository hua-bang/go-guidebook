package main

import "fmt"

func main() {

	for {
		fmt.Println("loop")
		break;
	}

	for j := 7; j <= 9; j++ {
		fmt.Println("loop", j)
	}

	for n := 0; n < 5; n++ {
		if n % 2 == 0 {
			continue
		}
		fmt.Println(n)
	}

	for w := 3; w < 10; w++ {
		if(w == 7) {
			break;
		}
		fmt.Println("w", w);
	}
}
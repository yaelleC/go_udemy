package main

import "fmt"

func main() {
	ints := make([]int, 11)

	for i := 0; i <= 10; i++ {
		ints[i] = i
	}

	for _, integer := range ints {
		if (integer % 2 == 0) {
			fmt.Println(integer, " is even")
		} else {
			fmt.Println(integer, " is odd")
		}
	}
}
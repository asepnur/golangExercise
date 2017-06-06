package main

import (
	"fmt"
)

func main() {
	x := 13 % 2
	if x == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}

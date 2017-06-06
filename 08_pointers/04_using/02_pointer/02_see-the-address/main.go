package main

import (
	"fmt"
)

func change(x *int) {
	fmt.Println("address & value in change function ")
	fmt.Println(x)
	*x = 0
}

func main() {
	z := 12
	fmt.Println("address & value in main ")
	fmt.Println(&z)
	fmt.Println(z)

	change(&z)
	// now z has changed
	fmt.Println(z)
}

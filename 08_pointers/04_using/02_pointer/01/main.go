package main

import (
	"fmt"
)

func change(x *int) {
	*x = 0
}
func main() {
	z := 12
	fmt.Println(z)
	fmt.Println(&z)

	change(&z)
	//z will change to 0
	fmt.Println(z)
}

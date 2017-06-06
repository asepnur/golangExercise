package main

import (
	"fmt"
)

func main() {
	var x int
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
	//and this is type of increment
	//
	fmt.Printf("type increment : %T \n", increment)
}

package main

import (
	"fmt"
)

func main() {
	a := 12
	fmt.Println("a : ", a)
	fmt.Println("address a : ", &a)

	b := &a
	fmt.Println("b : ", b)
	fmt.Println("address b : ", &b)
	//assign value or change value b with *b syntax
	*b = 50
	fmt.Println("a : ", a)
	fmt.Println("b : ", *b)

}

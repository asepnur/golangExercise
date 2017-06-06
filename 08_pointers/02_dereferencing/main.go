package main

import (
	"fmt"
)

func main() {
	a := 12
	fmt.Println("a : ", a)
	fmt.Println("mem address a :", &a)
	b := &a
	fmt.Println("b : ", b)
	//add * to see value that pointed b
	fmt.Println("reference value b : ", *b)

}

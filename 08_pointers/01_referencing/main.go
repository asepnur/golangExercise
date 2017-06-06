package main

import (
	"fmt"
)

func main() {
	a := 12
	fmt.Println("value a : ", a)
	fmt.Println("memory address a : ", &a)
	b := &a
	fmt.Println("value b : ", b)
}

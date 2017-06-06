package main

import (
	"fmt"
)

func wrapper() func() int {
	x := 12
	return func() int {
		x++
		fmt.Println("x : ", &x)
		return x
	}
}
func testing() func() string {
	x := "asepnur"
	return func() string {
		x = "muhammad"
		return x
	}
}
func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
	name := testing()
	fmt.Println(name())
	fmt.Println(name())
}

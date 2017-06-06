package main

import (
	"fmt"
)

func change(x int) {
	fmt.Println("x change function address")
	fmt.Printf("%p \n", &x)
	fmt.Println(&x)
	x = 0
}
func main() {
	x := 90
	println("x main address ")
	fmt.Printf("%p \n", &x)
	fmt.Println(&x)

	change(x)
	//still 90
	fmt.Println(x)
}

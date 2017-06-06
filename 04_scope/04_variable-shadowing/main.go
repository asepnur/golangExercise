package main

import (
	"fmt"
)

func myAge(x int) int {
	x = 12
	return x
}
func main() {
	//shadow variable
	//DONT do this, this is bad coding practise
	age := myAge(12)
	fmt.Println(age)
}

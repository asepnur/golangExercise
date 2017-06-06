package main

import "fmt"

var x = 12

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)
}

package main

import (
	"fmt"
)

func main() {
	x := 43
	fmt.Println("x - ", x)
	fmt.Println("Memory x - ", &x)
	fmt.Printf("x - %d", &x)
}

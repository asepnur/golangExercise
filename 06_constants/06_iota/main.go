package main

import (
	"fmt"
)

const (
	_ = iota
	a = iota
	b = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
}

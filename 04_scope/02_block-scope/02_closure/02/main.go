package main

import (
	"fmt"
)

//default value is zero
// you can also define like this
// var x = 0
var x int

func incSeq() int {
	x++
	return x
}

func main() {
	fmt.Println(incSeq())
	fmt.Println(incSeq())
}

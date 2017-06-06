package main

import (
	"fmt"
)

func change(x int) {
	x = 20
}

func main() {
	x := 12
	change(x)
	//x still 12
	fmt.Println(x)
}

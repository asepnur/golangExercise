package main

import (
	"fmt"
)

func main() {
	var i int
	for {
		fmt.Println(i)
		if i >= 10 {
			break
		}
		i++
	}
}

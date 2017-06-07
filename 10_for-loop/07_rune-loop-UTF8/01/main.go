package main

import (
	"fmt"
)

func main() {
	for i := 255; i <= 340; i++ {
		fmt.Println("i - ", i, "\t string ", string(i), " \tbyte ", []byte(string(i)))
		// for windows use this code if not print some character
		//fmt.Println("i - ", i, "\t string ", string(i), " \tbyte ", []int32(string(i)))
	}
}

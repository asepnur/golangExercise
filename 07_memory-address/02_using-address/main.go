package main

import (
	"fmt"
)

const metterToYard float64 = 1.09361

func main() {
	var x float64
	fmt.Print("Enter Metter : ")
	fmt.Scan(&x)
	result := x * metterToYard
	fmt.Println(result)
}

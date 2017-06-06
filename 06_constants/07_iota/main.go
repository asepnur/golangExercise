package main

import (
	"fmt"
)

const (
	_  = iota
	KB = iota << (1 * 10)
	MB = iota << (1 * 10)
	GB = iota << (1 * 10)
	TB = iota << (1 * 10)
)

func main() {
	fmt.Println("Decimal \t Binnary")
	fmt.Printf("%d\t%b\n", KB, KB)
	fmt.Printf("%d\t%b\n", MB, MB)
	fmt.Printf("%d\t%b\n", GB, GB)
	fmt.Printf("%d\t%b\n", TB, TB)
}

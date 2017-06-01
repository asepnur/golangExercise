package main

import "fmt"

func main() {
	for i := 1; i < 100; i++ {
		fmt.Printf("decimal : %d\tbinary : %b\thexadecimal : %#x\n", i, i, i)
	}
}

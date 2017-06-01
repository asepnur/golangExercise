package main

import "fmt"

func main() {
	for i := 60; i < 122; i++ {
		fmt.Printf("decimal : %d\tbinary : %b\thexa : %x\t utf-8 : %q \nt", i, i, i, i)
	}
}

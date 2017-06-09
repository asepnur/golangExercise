package main

import "fmt"

func main() {
	name := "asepnur"
	switch {
	case len(name) == 7:
		fmt.Println("Hello Asep Nur ")
	case name == "asep", name == "nur":
		fmt.Println("Hello asep or nur")
	case name == "iskandar":
		fmt.Println("Hello Iskandar")
	default:
		fmt.Println("Unknown")
	}
}

package main

import (
	"fmt"
)

func main() {
	switch "asepnur" {
	case "asep":
		fmt.Println("my name is asep")
	case "nur":
		fmt.Println("my name is nur")
	case "asepnur":
		fmt.Println("my name is asepnur")
	default:
		fmt.Println("unknow name")
	}
	/*
		no default fallthrough in golang
		fallthrough is optional in golang
		you can specify fallthrough by explicitly strting it
		break is not needed like in another language
	*/
}

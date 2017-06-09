package main

import "fmt"

func main() {
	switch "Asepnur" {
	case "Asepnur":
		fmt.Println("Hello Asepnur")
		fallthrough
	case "Muhammad":
		fmt.Println("Hello Muhammad")
		fallthrough
	case "Iskandar":
		fmt.Println("Hello Iskandar")
	case "Yusuf":
		fmt.Println("Hello Yusuf")
	default:
		fmt.Println("Unknown")
	}
}

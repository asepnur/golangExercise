package main

import "fmt"

func main() {
	switch "asep" {
	case "asep", "nur":
		fmt.Println("Halo asep or nur")
	case "muhammad", "iskandar":
		fmt.Println("Halo Muhammad or Iskandar")
	case "Yusuf", "abdullah":
		fmt.Println("Halo Yusuf or Abdullah")
	default:
		fmt.Println("Unknown")
	}

}

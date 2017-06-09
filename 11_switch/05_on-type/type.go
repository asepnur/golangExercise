package main

import "fmt"

type contact struct {
	name string
	age  int
}

//SwitchOnType dll
func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("integer")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}
func main() {
	var val int
	var str string
	SwitchOnType(val)
	SwitchOnType(str)
	var c = contact{
		"asepnur",
		22,
	}
	SwitchOnType(c)
	SwitchOnType(c.name)
	SwitchOnType(c.age)
}

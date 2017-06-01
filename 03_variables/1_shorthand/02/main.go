package main

import "fmt"

func main() {

	//this is how to declate varibale with short way
	//but this is only work under func
	a := 70
	b := "golang"
	c := 3.12
	d := false
	e := `Hello`
	f := 'A'

	//NOTICE!!!!
	// this is how to display value with default type of variable
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)

}

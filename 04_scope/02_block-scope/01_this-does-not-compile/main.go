package main

import (
	"fmt"
)

func main() {
	x := 12
	foo()
	foo2(x)

}

func foo() {
	//this is will be error because x is not package scope
	//add parameter to solve this probelem lets see another func in bellow function
	fmt.Println(x)
}

//it will work fine
//delete function foo to run this program
func foo2(x int) {
	fmt.Println(x)
}

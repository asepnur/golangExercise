package main

import "fmt"

func main() {
	name := "Asep Nur Muhammad"
	fmt.Println("name in main scope : ", name)
	{
		fmt.Println("name inner scope : ", name)
		age := 20
		fmt.Println("age in inner scope : ", age)
	}
	//you can't invoke var age in main scope
	//it will produce error
	//fmt.Println(age)
}

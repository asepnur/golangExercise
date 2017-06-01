package main

import "fmt"

//all 4 declaraion bellow are in package scope
var a = "this is stored in the variable a and in package scope"
var b, c string = "stored in b and in package scope", "stored in c and in package scope"
var d string

func main() {

	d = "stored in d" //assign d with value "stored in d"
	//function scope variable
	var e = 70
	f := 10
	g := "stored in g"
	h, i := "stored in h", "stored in i"
	j, k, l, m := 44.7, true, false, 'm'
	n := "stored in n" //use double quote
	o := 'o'           //use single quote
	p := `aku "suka" kamu `

	fmt.Println("a - ", a)
	fmt.Println("b - ", b)
	fmt.Println("c - ", c)
	fmt.Println("d - ", d)
	fmt.Println("e - ", e)
	fmt.Println("f - ", f)
	fmt.Println("g - ", g)
	fmt.Println("h - ", h)
	fmt.Println("i - ", i)
	fmt.Println("j - ", j)
	fmt.Println("k - ", k)
	fmt.Println("l - ", l)
	fmt.Println("m - ", m)
	fmt.Println("n - ", n)
	fmt.Println("o - ", o)
	fmt.Println("p - ", p)

}

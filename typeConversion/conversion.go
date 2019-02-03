package main

import "fmt"

type avocado int

var a int
var b avocado

func main() {
	a = 10
	b = 20
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	a = int(b) // assign the integer value of variable b (conversion) to the variable a
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}

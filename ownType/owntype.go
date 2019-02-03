package main

import "fmt"

type avocado int // create our own type AVOCADO with an underlying type int
var a avocado    // a is of type avocado

func main() {
	a = 10
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}

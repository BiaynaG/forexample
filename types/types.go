package main

import "fmt"

var y = 30

// Declaring the variable with the identifier "z" is of TYPE string
// Then this variable can only hold values of types string
var z = "Better safe than sorry"
var a = `Somebody said, "Better safe than sorry"`

func main() {
	fmt.Println(y)
	// Print the type of variable y in a new line
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}

package main

import "fmt"

// Declaration outside the function body
// VAR can be used at any place within the package; wide scope
// Declare  the variable "y" and assign the value 40
// declare & assign = initialization
var y = 40

// Declaring there is a VARIABLE with the IDENTIFIER "z", which is of TYPE int
// ASSIGNS the 0 value to integer "z"
// FALSE for booleans, 0 for integers, 0.0 for floats, "" for strings,
// NIL for pointers, functions, interfaces, slices, channels, and maps.
var z int

func main() {
	// Short declaration operator
	x := 10
	fmt.Println(x)

}

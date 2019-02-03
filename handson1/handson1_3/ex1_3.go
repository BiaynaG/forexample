package main

import "fmt"

var x = 42
var y = "James Bond"
var z = true

func main() {
	x = 42
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z) // %v gives the value in a default format
	fmt.Println(s)
}

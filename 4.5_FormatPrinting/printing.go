package main

import "fmt"

var y = 10

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)               // printing the type of the variable
	fmt.Printf("%b\n", y)               // printing the binary equivalent
	fmt.Printf("%x\n", y)               // printing the hexadecimal
	fmt.Printf("%#x\n", y)              // printing the hexadecimal with zero x in front
	fmt.Printf("%T\t%b\t%x\n", y, y, y) //priting multiple types on the same line

	// String printing
	s := fmt.Sprintf("%T\t%b\t%x\n", y, y, y)
	fmt.Println(s)
}

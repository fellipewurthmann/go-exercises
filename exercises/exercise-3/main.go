package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	// Using the code from exercise-2, at the package level scope, assign the following
	// values to the three variables
	// for x assign 42
	// for y assign "James Bond"
	// for z assign true
	// in func main use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN
	// the return value of type string using the short declaration operator to a VARIABLE
	// with the IDENTIFIER "s"

	s := fmt.Sprintf("%v\v%v\v%v", x, y, z)
	f := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	fmt.Println(s)
	fmt.Println()
	fmt.Println(f)
}

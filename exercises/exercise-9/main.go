package main

import "fmt"

func main() {
	// assigns an int to a variable, prints that int in decimal, binary and hex
	// shifts the bits of that int over 1 position to the left, and assigns that to a variable
	// prints that variable in decimal, binary and hex
	number := 42
	fmt.Printf("%d\t%b\t%U\n", number, number, number)

	number2 := number << 3
	fmt.Printf("%d\t%b\t%U", number2, number2, number2)
}

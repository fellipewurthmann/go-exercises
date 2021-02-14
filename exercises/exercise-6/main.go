package main

import (
	"fmt"
)

func main() {
	// Write a program that print a number in decimal, binary, hex and Unicode format
	number := 42

	fmt.Printf("%d\n", number)
	fmt.Printf("%b\n", number)
	fmt.Printf("%#X\n", number)
	fmt.Printf("%U\n", number)

}

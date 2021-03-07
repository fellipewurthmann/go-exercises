package main

import "fmt"

func main() {
	// Create a program that uses a switch statement with no switch expression specified
	switch {
	case true:
		fmt.Println("Is true! Pretty cool!!!")
	case false:
		fmt.Println("Is false! Not very cool!!!")
	default:
		fmt.Print("Neither true nor false...")
	}
}

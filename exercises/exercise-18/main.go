package main

import "fmt"

func main() {
	// Building on the previous hands-on exercise, create a program that uses "else if" and "else".
	name := "Salazar"
	if name == "James Bond" {
		fmt.Println("Bond, James Bond")
	} else if name == "Dr. No" {
		fmt.Println("Dr. No is here!")
	} else {
		fmt.Printf("I am %v", name)
	}
}

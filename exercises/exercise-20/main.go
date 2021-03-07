package main

import "fmt"

func main() {
	favSport := "Nothing"
	switch favSport {
	case "Football":
		fmt.Println("GOL!")
	case "Basketball":
		fmt.Println("POINT!")
	case "American Football":
		fmt.Println("TOUCHDOWN!")
	default:
		fmt.Println("I dont like sports.. :(")
	}
}

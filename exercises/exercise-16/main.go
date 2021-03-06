package main

import "fmt"

func main() {
	// Create a fot loop using this syntax: for {} ...
	// Have it print out the years you have been alive
	yearWasBorn := 1997
	for {
		fmt.Println(yearWasBorn)
		if yearWasBorn >= 2021 {
			break
		}
		yearWasBorn++
	}

}

package main

import "fmt"

func main() {
	// Using for sintax: for <condition> {} ...
	// Have it print out the years you have been alive
	yearWasBorn := 1997

	for yearWasBorn <= 2021 {
		fmt.Println(yearWasBorn)
		yearWasBorn++
	}
}

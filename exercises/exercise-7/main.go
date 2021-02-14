package main

import "fmt"

func main() {
	// Use OPERATORS to compare values, assign those values in variables and print
	numberOne := 42
	numberTwo := 43

	resultOne := numberOne == numberTwo
	resultTwo := numberOne <= numberTwo
	resultThree := numberOne >= numberTwo
	resultFour := numberOne != numberTwo
	resultFive := numberOne < numberTwo
	resultSix := numberOne > numberTwo

	fmt.Println(resultOne)
	fmt.Println(resultTwo)
	fmt.Println(resultThree)
	fmt.Println(resultFour)
	fmt.Println(resultFive)
	fmt.Println(resultSix)
}

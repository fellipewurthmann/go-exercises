package main

import "fmt"

const (
	untypedConst        = "Hello, world!"
	typedConst   string = "Hello, earth!"
)

func main() {
	// create untyped const and typed. then print those values
	fmt.Println(untypedConst)
	fmt.Println(typedConst)
	fmt.Printf("%T\t%T", untypedConst, typedConst)
}

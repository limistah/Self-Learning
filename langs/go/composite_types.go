package main

import (
	"fmt"
)

func main () {
	fmt.Println("Composite Types in go")

	// Bool -> Predefined constants true and false
	// String type -> Unicode chars written enclosed by double-qoutes


	// Array type
	// The length has to be specified
	var numbers = [4]int{1,2,3,4,}
	fmt.Printf("%T\n", numbers)

	// Can hold varying amount of data
	var slice = []string{}
	fmt.Printf("%T\n", (slice))

	// Map Type 
	// Similar to a Dictionay in python
	balances := map[string]float64{
		"USD": 2334.2,
	}
	fmt.Printf("%T\n", balances)

	// Struct type
	type Person struct{
		name string
		age int
	}

	var you Person
	fmt.Printf("%T\n", you)

	// Pointer -> 
	var x int = 2;
	ptr := &x
	fmt.Printf("ptr is of type %T with the value of %v\n", ptr, ptr)
}
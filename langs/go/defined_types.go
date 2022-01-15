package main

import "fmt"

func main() {
	type age int // Custom type with int as the base type
	type oldAge age // custom type referencing age, still, int, not age as the base type
	type veryOldAge age // custom type referencing oldAge, int is the base type


	// We can attach methods to new type
	// Readablity e.g type usd float64

	type speed uint
	var s1 speed = 10
	var s2 speed = 10

	fmt.Println(s1 - s2)
}
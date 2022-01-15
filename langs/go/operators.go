package main

import "fmt"

func main () {
	a, b := 4,2

	r := (a+b) / (a-b) * 2

	fmt.Println(r)

	r = 9 % a
	fmt.Println(r)

	// Assignment operators


	a, b = 2,3
	a +=b // increment assignment
	fmt.Println(a)
	// decremtn assignment
	b -= 2
	fmt.Println(b)

	// multiplication assignmend
	b *= 10
	fmt.Println(b)

	// division assignment
	b /= 5
	fmt.Println(b)

	// modulus assignment

	a %= 3
	fmt.Println(a)

	// These assignments are statements and not operators, the below example is illegal in go
	// fmt.Println(5 - a--)


	// Comparison operators
	a,b = 5, 10
	fmt.Println(a == b) // false
	fmt.Println(a != b) // true
	fmt.Println(a > b, a >= b) // false, false
	fmt.Println(a < b, a <= b) // true, true

	// Logical operators
	a, b = 5, 10
	fmt.Println(a < 1 && b == 10) // false
	fmt.Println(a == 5 || b == 100) // true
	fmt.Println(!(a == 1))

}
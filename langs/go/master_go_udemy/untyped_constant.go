package main

import (
	"fmt"
)

func main () {

	//UNTYPED CONTSTANTS
	const a float64 = 5.1 // typed constant

	const b = 6.7 //untyped constant

	const d = 5 > 10

	fmt.Println(d);

	const x int = 5
	// We can't multiple a typed int constant with a float
	// Considering go rules.
	// const y float64 = 2.2 *x;

	const xx = 5
	// This is possible because untyped constants does not obey go strict type rules
	const yy  = 2.2 * 5

	var i int = x // changes x to an int
	var j float64 = xx // var j float64 = float64(x)
	fmt.Println(i, j)

	const r = 5 // an untyped int constant, because 5 is of type
	var rr = r // Converts the literal of 5 to a typed integer
	fmt.Printf("r:%T, rr:%T \n",r, rr)
}
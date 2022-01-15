package main

import (
	"fmt"
)

func main () {
	price, inStock := 100, true

	if price > 80 {
		fmt.Println("Too Expensive")
	}

	if price <= 100 && inStock {
		fmt.Println("Buy it!")
	}

	// There is no truthfulness of a variable
	// if price {
	// 	fmt.Println("Somehthing")
	// }
	if price < 100 {
		fmt.Println("it's cheap")
	} else {
		fmt.Println("It's expensive")
	}

	age := 50
	if age >= 0 && age < 18 {
		fmt.Printf("You cannot vote! Please return on %d years !\n", 18-age)
	} else if age ==18 {
		fmt.Printf("This is your first vote!")
	} else if age > 18 && age <= 100 {
		fmt.Printf("Please vote, it's important!")
	} else {
		fmt.Printf("Invalid age!")
	}
}
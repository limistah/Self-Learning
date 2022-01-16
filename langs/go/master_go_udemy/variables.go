package main

import "fmt"

// import

func main () {

	var age int = 30
	// Inferred type
	var name = "Dan"

	fmt.Println("Your name is:", name, "and age is", age)

	// Short declaration and reassignable
	new_name := "John"
	// blank identifier makes the variable go unsued
	_ = new_name
	// Can not be use to reassign a variable
	// name := "John"

	// declare many variables on a line
	car, cost := "Audi", 2018
	fmt.Println(car, cost)
	// You can't reassign existing variables
	// car, cost := "BMW", 2018
	// A new variable name must exist to reassign to old ones
	car, year := "BMW", 2018
	fmt.Println(car, year)

	var opened = false
	opened, file := true, "file.txt"
	fmt.Println(opened, file)

	var (
		salary float64
		firstName string
		gender bool
	)
	fmt.Println(salary, firstName, gender)

	var i, j int
	i, j  = 5, 8
	_, _ = i, j


}
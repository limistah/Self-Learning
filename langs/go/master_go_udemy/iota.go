package main

import (
	"fmt"
)

func main () {
	fmt.Println("This is iota")

	// Incrementing integers for a contant 
	const (
		c11 = iota
		c22 = iota
		c33 = iota
	)

	fmt.Println(c11, c22, c33)

	const (
		c21 = iota
		c24
		c23
	)
	// increments by 1 automatically same as using iota for each item
	fmt.Println(c21, c24, c23)

	// initializing with a step
	const (
		a = (iota + 1) * 2 // 0
		b
		c
	)
	fmt.Println(a, b, c)

	// skip a value
	const (
		x = -(iota + 2) // -2
		_ // Skips -3
		y
		z
	)
	fmt.Println(x,y,z)

}
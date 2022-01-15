package main

import (
	"fmt"
)

func main () {
	// Untyped constant
	const days int = 7

	var i int
	fmt.Println(i)

	const pi float64 = 3.14
	const secondsInHour = 3600

	duration := 234
	fmt.Printf("Duration in second: %v\n", duration * secondsInHour)

	// x,y:=5,0
	// fmt.Println(x/y)

	const a,b = 5,0

	const n, m int = 4,5

	// Multiple declaration
	const (
		min1 = -500
		min2
		min3
	)

	fmt.Println(min1, min2, min3) // -> -500 -500 -500
	
	const temp int = 100
	// 1. Can not change a constant
	// temp = 50

	// 2. You can not initiate a contant at run time
	// const power = math.Pow(2,3)

	// 3. You can not use a variable to intialize a constant
	// t :=5
	// const tc = t



}
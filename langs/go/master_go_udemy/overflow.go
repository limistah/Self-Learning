package main

import (
	"fmt"
	"math"
)

func main () {
	var x uint8 = 255
	x++ // would overflow
	// Go would cath this at runtime
	fmt.Println(x) // -> 0

	// Go would catch this at compile time
	// a := int8(255 + 1)

	var b int8 = 127
	// The below arithmetic would overflow the b variable, and since it is going above the the maximum value, go sets the variable to the minimum possible value
	fmt.Printf("%d\n", b+1) // -128


	var c int8 = -128
	// The below arithmetic would overflow the b variable, and since it is going below the minimum value, go sets the variable to the maximum possible value
	fmt.Printf("%d\n", c-1) // 127

	f := float32(math.MaxFloat32)

	fmt.Println(f)
	f = f * 1.2 // Overflows to infinity
	fmt.Println(f)

	// Use the big variable if your application requires big numbers
}
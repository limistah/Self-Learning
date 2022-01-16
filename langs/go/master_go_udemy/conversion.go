package main

import (
	"fmt"
	"strconv"
)

func main () {
	var x = 3 // type int
	var y = 3.1 // float64 type
	// x = x * y
	x = x * int(y)
	fmt.Println(x, y)
	fmt.Printf("x: %T, y: %T", x, y)

	x = int(float64(x) * y)
	fmt.Println((x))

	y = float64(x) * y
	fmt.Println(y)

	a := 5
	fmt.Printf("%T\n", a)

	var b int64 = 2
	_ = b
	// Explicit conversions are required when different numeric types are mixed in an assignment or numeric operations
	// a = b // -> Would fail b is of type int64 and can not be applied to a of type int
	a = int(b)


	// Numbers to strings
	s := string(99)
	fmt.Println(s)

	// converting a float to a string
	// the string function won't work, we have to use the Sprintf here
	myStr := fmt.Sprintf("%f", 44.2)
	fmt.Println(myStr)

	myStr1 := fmt.Sprintf("%d", 34234)
	fmt.Println(myStr1)

	fmt.Println(string(34234))

	// Strings to Number conversiona
	s1 := "3.123"
	fmt.Printf("%T\n", s1)
	var f1, err = strconv.ParseFloat(s1, 64)
	_ = err
	fmt.Println(f1 * 3.4)

}
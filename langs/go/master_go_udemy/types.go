package main

import (
	"fmt"
)

func main () {
	fmt.Println("Types in go")

	// Numeric types in go
	// int8, int16, int64
	// unint8, uint16, uint32, uint64 unsigned positive integers.
	// unint is an alias for unint32 or unint64 based on platform.
	// int is an alias for int32 or int64 based on platform
	// float32, float64: zero before the decimal point separator can be omitted (-.5 -3 -0 1.4)
	// complex64 complex128
	// byte alias for uint8
	// rune alias for int32

	// INT Type
	var i1 int8 = -128
	fmt.Printf("%T\n", i1)

	var i2 uint16 = 65535
	var i3 uint32 = 4294967295
	var i4 uint64 = 18446744073709551615
	var i5 int16 = 32767
	var i6 int16 = -32768
	var i7 int32 = 2147483647
	var i8 int32 = -2147483648
	var i9 int64 = 9223372036854775807
	var i10 int64 = -9223372036854775808
	var _, _, _, _, _, _, _ = i4, i5, i6, i7, i8, i9, i10
	fmt.Printf("%T\n", i2)
	fmt.Printf("%T\n", i3)

	// float type
	var f1, f2, f3 float64 = 1.1, - 0.2, 5.0
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	// RUNE TYPE
	var r rune = 'f' // use to store the hexadecimal values of a string
	fmt.Printf("%T\n", r)
	fmt.Println(r)
	fmt.Printf("%x\n", r) // Prints the hexadecimal ascii code for the string

	var b bool = true
	fmt.Println(b);

	
}
// Pointers allow you to share data between functions.
// However, when sharing data between functions and goroutines,
// you should be extra careful with race condition issues.
// Pointers are also very handy when you want to tell the difference
// between the zero value of a variable and a value that is not set (nil).
//This is particularly useful with structures because pointers (and therefore pointers to structures,
//which are fully covered in the next chapter), can have the nil value,
//which means that you can compare a pointer to a structure with the nil value,
//which is not allowed for normal structure variables.
//Having support for pointers and, more specifically,
//pointers to structures allows Go to support data structures such as linked lists and binary trees,
//which are widely used in computer science.
//Therefore, you are allowed to define a structure field of a Node structure as Next *Node,
//which is a pointer to another Node structure.
//Without pointers, this would have been difficult to implement and may be too slow.

package main

import "fmt"

type aStructure struct {
	field1 complex128
	field2 int
}

func processPointer(x *float64) {
	*x = *x * *x
}

//In order to return the memory address of a regular variable, you need to use & (&temp).
func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func main() {
	var f float64 = 12.123
	fmt.Println("Memory address of f:", &f)
	// Pointer to f
	fP := &f
	fmt.Println("Memory address of f:", fP)
	fmt.Println("Value of f:", *fP)
	// The value of f changes processPointer(fP)
	fmt.Printf("Value of f: %.2f\n", f)

	// The value of f does not change
	x := returnPointer(f)
	fmt.Printf("Value of x: %.2f\n", *x)

	// The value of f does not change
	xx := bothPointers(fP)
	fmt.Printf("Value of xx: %.2f\n", *xx)

	// Check for empty structure
	var k *int

	// This is nil because currently k points to nowhere
	fmt.Println(k)
	// Therefore you are allowed to do this:
	if k == nil {
		k = new(int)
	}

	fmt.Printf("%+v\n", k)
	if k != nil {
		fmt.Println("k is not nil!")
	}
}

package main

import "fmt"

func main () {
	name := "Aleem"
	fmt.Println("Hello, playground", name)

	a, b := 4,6

	fmt.Println("Sum:", a+b, "Meaa Value:", (a+b) /2)

	fmt.Printf("Your age is %d", 21)

	fmt.Printf("x is %d, y is %2f", 5, 6.8)

	fmt.Printf("Your age is %d\n", 21)

	fmt.Println("He says: \"Hello Go!\"")

	figure := "Circle"
	radius := 5
	pi := 3.14159
	_, _ = figure, pi

	// %d for integer
	fmt.Printf("Radius is %d\n", radius)
	fmt.Printf("Radius is %+d\n", radius)
	fmt.Printf("He says: %q\n", figure)


	fmt.Printf("Pi constant is %f\n", pi)

	fmt.Printf("The diameter of a %s with a Radius of %d is %f\n", figure, radius, float64(radius)*2*pi)

	// %T for type
	fmt.Printf("figure is of type %T\n", figure)
	fmt.Printf("radius is of type %T\n", radius)
	fmt.Printf("pi is of type %T\n", pi)

	// %t for boolean
	closed := true
	fmt.Printf("File closed: %t\n", closed)

	// %b -> base 2
	fmt.Printf("%08b\n", 50)


	// %x -> base 16
	fmt.Printf("%x\n", 100)
}
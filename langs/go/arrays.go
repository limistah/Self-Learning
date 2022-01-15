package main

import "fmt"


func main () {
	fmt.Printf("Array")

	var numbers [4]int

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)
	
	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [3]int{-10, 1, 100}
	fmt.Printf("%#v\n", a2)

	var a3 = [4]string{"Dan", "Diana", "Paul", "John"}
	fmt.Printf("%#v\n", a3)

	a4 := [4]string{"x", "y"}
	fmt.Printf("%#v\n", a4)

	// ellipsis operator
	a5 := [...]int{1,2,3,4,55,56,7,8,9,-10}
	fmt.Printf("%#v\n", a5)
	fmt.Printf("the length of a5 is %d\n", len(a5))


	a6 := [...]int{
		1,
		2,
		3,
		4,
		5, // this comman is required for multiline declaraion
	}
	_ = a6

	// accessing an array
	numbers = [4]int{0,1,3}
	numbers[0] = 7
	fmt.Printf("%#v\n", numbers)
	// numbers[5] = 100 // error no index 5

	// Looping an array
	for i,v := range numbers {
		fmt.Println("index:", i, " value: ", v)
	}
	for i :=  0; i<len(numbers); i++ {
		fmt.Println("index:", i, " value: ", numbers[i])
	}

	// Multidimensional array
	balances := [2][3]int{
		{5,6,7},
		[3]int{8, 9, 10},
	}
	fmt.Printf("%#v\n", balances)

	m := [3]int{1,2,3}
	n := m // the array at m is copied over to n and not passed as a reference
	fmt.Println("n is equal to m: ", n ==m) 

	// Arrays keyed elements
	grades := [3]int{
		1: 10,
		0: 5,
		2: 7,
	}

	fmt.Println(grades)

	accounts := [3]int{2:50}
	fmt.Println(accounts)

	names := [...]string{
		5: "Dan",
	}

	fmt.Println(names, len(names))

	cities := [...]string{
		5: "Paris",
		"London", // index 6
		1: "NYC",
	}

	fmt.Printf("%#v\n", cities)

	weekend := [7]bool{5: true, 6: true}
	fmt.Println(weekend)
}
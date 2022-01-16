package main

import (
	"fmt"
	"unsafe"
)


func main () {
	// A slice can shrink and grow at runtime
	// While an array has fixed size at compile time
	var cities []string
	fmt.Println("cities is equal to nil: ", cities == nil)
	fmt.Printf("cities %#v\n", cities)

	// cities[0] = "London" // -> Would throw index out of range
	numbers := []int{2,3,4,5}
	fmt.Println(numbers)

	nums := make([]int, 2)
	fmt.Printf("%#v\n", nums)

	type names []string
	friends := names{"Dan", "Maria"}
	fmt.Println(friends)

	// Getting from a slice index, returns the value and not a reference
	myFriend := friends[0]
	fmt.Println("My best friend is ", myFriend)

	// Reassigning at a slice index
	friends[0] = "Gabriel"
	fmt.Println("My best friend is ", friends[0])

	for index, value := range numbers {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	// Comparing slices
	var n []int
	fmt.Println(n == nil)

	m := []int{}
	fmt.Println(m == nil)

	a, b := []int{1,2,3}, []int{1,2,3}

	// fmt.Println(a == b) // would throw
    var eq bool = true

	if (len(a) != len(b)) {
		eq = false
	} else {
		for i, valueA := range a {
			if valueA != b[i] {
				eq = false
				break
			}
		}
	}

	if (eq) {
		fmt.Println("a and b are equal")
	} else {
		fmt.Println("a and b are not equal")
	}

	// Appending to slice
	numbers = []int{2,3}
	numbers = append(numbers, 1)
	fmt.Println(numbers)
	numbers = append(numbers, 20, 30, 40)
	fmt.Println(numbers)
	// overwrite the element of a slice at an index
	numbers = append(numbers[:4], 200)
	fmt.Println(numbers)

	n = []int{100,200}
	numbers = append(numbers, n...)
	fmt.Println(numbers)

	// Copy slices
	src := []int{100, 20, 30}
	dst := make([]int, len(src))
	nn := copy(dst, src)
	fmt.Println(dst, dst, nn)

	/**
	Slice experessions, slice backing(underlying array)
	When creating a slice behind the scenes Go creates a hidden array called Backing Array
	The backing array stores the elements not the slice
	Go implements a slice as a data structure called slice hearder.

	A slice header contains 3 fields:
	- the address of the backing array(pointer)
	- the length of the slice. len() returns it.
	- the capacity of the slice cap() returns it.

	Slice Header is the runtime representation of a slice
	A nil slice doesnt havve a backing array

	**/
	s1 := []int{10,20,30,40,50}
	// s3 selects from 0, and picks 2 items
	// s4 selects from 1 and picks 2 items
	s3, s4 := s1[0:2], s1[1:3]

	s3[1] = 600
	fmt.Println(s1)
	fmt.Println(s4)

	arr1 := [4]int{10, 20, 30,40}
	// Slice holds a reference to all the slices
	slice1, slice2 := arr1[0:2], arr1[1:3]
	
	// Create a variable that holds the same value as the first item
	el1 := arr1[1]
	// This modifies all the slices as well, since slices are references
	arr1[1] = 2

	
	fmt.Println(arr1, slice1, slice2, el1)

	// To copy to a slice, use append
	cars := []string{"Ford", "Audi", "Honda", "Range Rover"}
	newCars := []string{}
	// Copy by value the slice of cars from 0 to 2
	newCars = append(newCars, cars[0:2]...)

	// Modify the first element of cars to Nissan, which wouldn't update the value of newCars
	cars[0] = "Nissan"
	fmt.Println(cars, newCars)

	s1 = []int{10,20,30,40,50}
	newSlice := s1[0:3]
	fmt.Println(len(newSlice), cap(newSlice))
	newSlice = s1[2:5]
	fmt.Println(len(newSlice), cap(newSlice))

	fmt.Printf("%p\n", &s1)

	fmt.Printf("%p %p \n", &s1, &newSlice)
	newSlice[0] = 1000

	fmt.Println("s1: ", s1)

	arr := [5]int{1,2,3,4,5}
	s := []int{1,2,3,4,5}

	fmt.Printf("array's size in bytes: %d \n", unsafe.Sizeof(arr))
	fmt.Printf("slice's size in bytes: %d \n", unsafe.Sizeof(s))

	// Append internals
	var numss []int
	fmt.Printf("%#v\n", numss)

	fmt.Printf("Length: %d, Capacity: %d \n", len(numss), cap(numss)) // 0, 0

	numss = append(numss, 1,2)
	fmt.Printf("Length: %d, Capacity: %d \n", len(numss), cap(numss)) // 2, 2
	numss = append(numss, 3)
	// append adds extra space to the underlying array to avoid recreating a new backing array everytime
	fmt.Printf("Length: %d, Capacity: %d \n", len(numss), cap(numss)) // 3, 4

	numss = append(numss, 4)
	fmt.Printf("Length: %d, Capacity: %d \n", len(numss), cap(numss)) // 4, 4


	numss = append(numss, 100)
	fmt.Printf("Length: %d, Capacity: %d \n", len(numss), cap(numss)) // 5, 8

	fmt.Println(numss)
	fmt.Printf("Length: %d, Capacity: %d \n", len(numss), cap(numss)) // 5, 8


	// strange slice append
	letters := []string{"A", "B", "C", "D", "E", "F"}
	letters = append(letters[0:1], "X", "Y") // pick the first item, and append 2 items to the new slice, making 3
	// Below outputs 3, 6
	// The first array is the underlying array which is 6 elements
	// Because we are storing the new slice to the same slice variable,
	// go just updates the underlying slice
	// and update the length of the new slice
	fmt.Printf("Length: %d, Capacity: %d \n", len(letters), cap(letters)) 
	// We can't use the index expression to get an item from the slice
	// fmt.Println(letters[3]) // would through index out of bound error
	// We can use the slice operator to get the remaining values after the known length
	fmt.Println(letters[3:6])

}
package main

import "fmt"

func main() {
	// map
	var employees map[string]string

	fmt.Printf("%#v\n", employees)
	fmt.Printf("Number of pairs:%#d\n", len(employees))

	fmt.Printf("The value for key %q is %q\n", "Dan", employees["Dan"])

	var accounts map[string]float64
	fmt.Printf("%v\n", accounts["0x323"])

	var m1 map[[5]int]string
	_ = m1
	//employees["Dan"] = "Programmer"

	people := map[string]float64{}

	people["John"] = 21.4
	people["Mary"] = 10
	fmt.Println(people)

	// Using make to declare a map
	map1 := make(map[int]int)
	map1[4] = 8

	balances := map[string]float64{
		"USD": 323.11,
		"EUR": 322.1,
		"CHF": 3243.1,
	}
	fmt.Println(balances)

	m := map[int]int{1: 10, 2: 20, 3: 30}
	_ = m

	balances["USD"] = 500.2
	balances["GBP"] = 100
	fmt.Println(balances)

	fmt.Println(balances["RON"]) // does not exist, therefore the zero value of the map value type is returned

	balances["RON"] = 100
	v, ok := balances["RON"]

	if !ok {
		fmt.Println("RON Key does not exists")
	} else {
		fmt.Println("The RON balance is: ", v)
	}

	// Looping over map
	// Not documented, because maps have been optimized for reads and the keys change often
	for k, v := range balances {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v)
	}

	// Deleting a map
	delete(balances, "RON")
	fmt.Println(balances)

	// Comparing Maps, though it is not possible to compare maps
	a := map[string]string{"A": "X"}
	b := map[string]string{"A": "Y"}

	//fmt.Println(a == b) // Throws

	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)

	if s1 == s2 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}

	// Cloning and duplicating a map
	/**
	Go creates a Map Header in memory
	the map references the internal data structure
	The map contains only the memory address of the map header
	The key value pairs are not stored in the map directly but into the memory address of the map header
	When copying a map, we are only copying a the reference to the map
	And not the map header that holds the key value pairs
	*/
	friends := map[string]int{"Dan": 40, "Maria": 25}
	neighbours := friends

	friends["Dan"] = 50

	fmt.Println("Friends", friends)
	fmt.Println("Neighbours", neighbours)
	// To copy a map, initialize it and use a for loop
	people2 := make(map[string]int)
	for k, v := range friends {
		people2[k] = v
	}
	people2["Anne"] = 18 // modify just people, and not friends
	fmt.Printf("%#v ----------- %#v", people2, friends)
}

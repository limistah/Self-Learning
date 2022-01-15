package main

import "fmt"

func main () {
	var a = 4
	var b = 3.2

	a = int(b)
	fmt.Println(a, b)
}
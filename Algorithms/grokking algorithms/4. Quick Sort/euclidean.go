package main

import "fmt"

func euclidean(a int, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	remainder := a % b
	if remainder == 0 {
		return b
	}
	return euclidean(b, remainder)
}

func main() {
	gcd := euclidean(270, 192)
	fmt.Println(gcd)
}

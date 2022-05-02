package main

import "fmt"

func funcRet(i int) func(int) int {
	if i < 0 {
		return func(k int) int {
			k = -k
			return k + k
		}
	}
	return func(k int) int {
		return k * k
	}
}

func main() {
	ret := funcRet(-1)
	fmt.Printf("ret %d\n", ret(10))
}

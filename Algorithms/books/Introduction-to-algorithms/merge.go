package main

import (
	"fmt"
	"strings"
)

func merge(A []int32, p int32, q int32, r int32) []int32 {
	n1 := q - p + 1
	n2 := r - q
	L := make([]int32, n1+1)
	R := make([]int32, n2+1)

	for i := int32(0); i < n1; i++ {
		L[i] = A[p+i-1]
	}

	for j := int32(0); j < n2; j++ {
		R[j] = A[q+j]
	}

	i := 0
	j := 0
	s := make([]int32, r-(p-1))
	//
	start := p - 1
	step := int32(1)
	for i := range s {
		s[i] = start
		start += step
	}
	//fmt.Println(A, p, r, L, R, s, i, j)

	//fmt.Println(s, p, r)
	//
	for _, key := range s {
		fmt.Println(i, j, len(R))

		fmt.Println(L[i], R[j], L, R, key, s, A)

		if len(R) > j {
			if L[i] <= R[j] {
				A[key] = L[i]
				i = i + 1
			} else {
				A[key] = R[j]
				j = j + 1
			}
		} else {
			A[key] = L[i]
			i = i + 1
		}

	}
	fmt.Println(strings.Repeat("---", 5))
	return A
}

func mergeSort(A []int32, p int32, r int32) []int32 {
	if p < r {
		q := (p + r) / 2
		mergeSort(A, p, q)
		mergeSort(A, q+1, r)
		merge(A, p, q, r)
	}
	return A
}

func main() {
	fmt.Println("Merge Algorithms")
	slice := []int32{31, 41, 59, 26, 42, 58, 4, 1, 3, 0}
	sorted := mergeSort(slice, 1, int32(len(slice))-1)
	fmt.Println(sorted)
	//newSlice := []int32{}
	//sorted := merge(append(newSlice, slice[0:]...), 0, 0, int32(len(newSlice)))
	//fmt.Printf("%v\n", sorted)
}

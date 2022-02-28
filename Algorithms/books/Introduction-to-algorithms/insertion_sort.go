package main

import (
	"fmt"
)

func insertionSort(A []int32) []int32 {
	for i := 1; i < len(A); i++ {
		prevItemIndex := i - 1
		currentItem := A[i]
		for prevItemIndex > -1 && A[prevItemIndex] > currentItem {
			A[prevItemIndex+1] = A[prevItemIndex]
			prevItemIndex = prevItemIndex - 1
		}
		A[prevItemIndex+1] = currentItem
	}
	return A
}

func reversedInsertionSort(A []int32) []int32 {
	for i := 1; i < len(A); i++ {
		key := A[i]
		j := i - 1
		for j > -1 && key > A[j] {
			A[j+1] = A[j]
			j = j - 1
		}
		A[j+1] = key
	}
	return A
}

func main() {
	fmt.Println("Insertion sort algorithm")
	slice := []int32{31, 41, 59, 26, 41, 58, 4, 1, 3, 0}
	newSlice := []int32{}
	//sorted := insertionSort(append(newSlice, slice[0:]...))
	revsorted := reversedInsertionSort(append(newSlice, slice[0:]...))
	fmt.Printf("Sorted %v into %v \n", slice, revsorted)
}

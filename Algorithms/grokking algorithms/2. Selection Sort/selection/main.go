package main

import (
	"fmt"
	"math"
)

func findSmallestIndex(items []int) int {
	smallestIndex := 0
	smallestValue := math.MaxInt64 // Assign to the highest possible number
	for i, v := range items {
		if v <= smallestValue {
			smallestIndex = i
			smallestValue = v
		}
	}
	return smallestIndex
}

func selectionSort(items []int) []int {
	finalSlice := make([]int, 0)
	for _, _ = range items {
		smallestIndex := findSmallestIndex(items)
		finalSlice = append(finalSlice, items[smallestIndex])
		items = append(items[:smallestIndex], items[smallestIndex+1:]...)
	}
	return finalSlice
}

func main() {
	fmt.Println(selectionSort([]int{1, 5, 3, 6, 2, 10, 4, 8, 7}))
}

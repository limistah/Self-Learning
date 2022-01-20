package main

import "fmt"

/**
Quick sort uses a Divide and Conquer approach
Which indirectly uses recursion
For recursion, we need the base case and the iterative case.

Quick sort
- Don't sort an array with just an item
- Don't sort an array with one element

- Pick an item from the array (called pivot)
- Store all items greater than the pivot in the array excluding the pivot in a variable
- store all items lesser than the pivot in the array excluding the pivot in a variable
- Recursively call quicksort with a merge of lesserItems + pivot + greaterItems
*/

func quickSort(items []int) []int {
	if len(items) < 2 {
		return items
	}
	pivotIndex := len(items) / 2
	pivot := items[pivotIndex]
	leftItems := make([]int, 0)
	rightItems := make([]int, 0)
	for _, val := range append(items[:pivotIndex], items[pivotIndex+1:]...) {
		if val <= pivot {
			leftItems = append(leftItems, val)
		}
		if val > pivot {
			rightItems = append(rightItems, val)
		}
	}
	return append(append(quickSort(leftItems), pivot), quickSort(rightItems)...)
}

func main() {
	fmt.Println("Quick Sort Here")
	fmt.Println(quickSort([]int{1, 10, 50, 90, 100, 100, 2, 4, 4, 5, 3}))
}

package main

import "fmt"

func sum(arr []int, aggregate int) int {
	if len(arr) == 0 {
		return aggregate
	}
	return sum(arr[1:], arr[0:1][0]+aggregate)
}

func count(arr []string) int {
	if len(arr) == 0 {
		return 0
	} else {
		return count(arr[1:]) + 1
	}
}

func max(arr []int, currentMax int) int {
	if len(arr) == 0 {
		return currentMax
	}
	itm := arr[0:1][0]
	if itm > currentMax {
		currentMax = itm
	}
	return max(arr[1:], currentMax)
}

func bSearch(arr []int, needle int, index int) (int, string) {
	if len(arr) == 1 {
		return index, ""
	}
	slc := make([]int, len(arr))
	midIndex := (len(arr) - 1) / 2
	midItem := arr[midIndex]
	if needle == midItem {
		return midIndex, ""
	} else if needle > midItem {
		slc = arr[midIndex+1:]
	} else {
		slc = arr[:midIndex]
	}
	return bSearch(slc, needle, midIndex)
}

func binarySearch(arr []int, needle int) (int, string) {
	return bSearch(arr, needle, 0)
}

func main() {
	scores := []int{10, 20, 30, 100, 200}
	fmt.Println(max(scores, 0))
	fmt.Println(count([]string{"Aleem", "Isiaka", "Aremu"}))

	fmt.Println(binarySearch(scores, 200))
	// Binary search base case when array is zero
	//
}

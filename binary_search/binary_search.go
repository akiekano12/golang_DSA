package main

import "fmt"

func binarySearch(needle int, hayStack []int) bool {
	low := 0
	high := len(hayStack) - 1

	for low <= high {
		median := (low + high) / 2

		if hayStack[median] != needle {
			low = median + 1
		} else {
			high = median -1
		}
	}

	if low == len(hayStack) || hayStack[low] != needle {
		return false
	}

	return true
}

func main() {
	items := []int{1,2,3,42,14,45,4,9}
	fmt.Println(binarySearch(63, items))
}

// Binary Search is a data search technique by 
// repeatedly dividing a portion of the amount 
// of data sought to reduce the search location 
// to one data.

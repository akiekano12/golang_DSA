package main

import "fmt"

func interpolationSearch(array []int, key int) int {
	min, max := array[0], array[len(array) - 1]
	low, high := 0, len(array) - 1

	for {
		if key < min {
			return low
		}

		if key > max {
			return high + 1
		}

		// guess of the loaction
		var guess int
		if high == low {
			guess = high
		} else {
			size := high - low
			offset := int(float64(size - 1) * (float64(key - min) / float64(max - min)))
			guess = low + offset
		}

		// if found
		if array[guess] == key {
			// scan backwards for start of val range
			for guess > 0 && array[guess - 1] == key {
				guess--
			}

			return guess
		}

		// if we guessed to high, guess lower or vice versa
		if array[guess] > key {
			high = guess - 1 
			max = array[high]
		} else {
			low = guess + 1
			min = array[low]
		}
	}
}

func main() {
	items := []int{1,2,3,4,13,31,42,53,94}
	fmt.Println(interpolationSearch(items, 31))
}

// searching algorithm for sorted arrays. It is based
// on the idea that if an element is located in some 
// portion of an array, then it is likely to be in 
// the middle of that portion.

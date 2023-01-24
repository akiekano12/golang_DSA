package main

import "fmt"

func linierSearch(dataList []int, key int) (bool, int) {
	for i, item := range dataList {
		if item == key {
			return true, i
		}
	}
	return false, 0
}

func main() {
	items := []int{95,12,41,45,6,2,12,3}
	fmt.Println(linierSearch(items, 41))
}

// Linear Search is defined as a sequential search algorithm
// that starts at one end and goes through each element of
// a list until the desired element is found, otherwise the 
// search continues till the end of the data set. It is 
// the easiest searching algorithm

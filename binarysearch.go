/*package main

import "fmt"

func main() {
	array := []int{10, 20, 30, 40, 50, 60}
	result := binarysearch(array, 10)
	if result == -1 {
		fmt.Println("key does not exist")
	} else {
		fmt.Println("key found at index:", result)
	}
}

// returns the index
// return -1 if it does not exist
func binarysearch(array []int, key int) int {
	first := 0
	last := len(array) - 1
	for first <= last {
		mid := first + (last-first)/2
		if key < array[mid] {
			last = mid - 1
		} else if key > array[mid] {
			first = mid + 1
		} else {
			return mid
		}
	}
	return -1
}*/

package main

import "fmt"

func main() {
	array := []int{10, 20, 30, 40, 50, 60}
	result := binarysearch(array, 70)
	if result == -1 {
		fmt.Println("key does not exist")
	} else {
		fmt.Println("key found at index:", result)
	}
}

func binarysearch(array []int, key int) int {
	first := 0
	last := len(array) - 1

	// find wheather the array is sorted in acsending or descending order
	isAsc := array[first] < array[last]

	for first <= last {
		mid := first + (last-first)/2
		if array[mid] == key {
			return mid
		}
		if isAsc {
			if key < array[mid] {
				last = mid - 1
			} else {
				first = mid + 1
			}
		} else {
			if key > array[mid] {
				last = mid - 1
			} else {
				first = mid + 1
			}
		}

	}
	return -1 //if key does not exist array
}

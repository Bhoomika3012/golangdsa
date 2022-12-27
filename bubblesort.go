package main

import "fmt"

func main() {
	arr := []int{10, 6, 2, 4, 8}
	fmt.Println(arr)
	bubblesort(arr)
	fmt.Println(arr)
}

func bubblesort(arr []int) {
	n := len(arr)
	var isSorted bool
	for pass := n - 1; pass >= 0; pass-- {
		isSorted = true

		for i := 0; i < pass; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				isSorted = false
			}
		}
		if isSorted {
			break
		}
	}
}

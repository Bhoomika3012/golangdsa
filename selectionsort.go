package main

import "fmt"

func main() {
	arr := []int{40, 90, 10, 60, -5, -10}
	fmt.Println(arr)
	selectionsort(arr)
	fmt.Println(arr)
}
func selectionsort(arr []int) {
	fmt.Println("")
	var pos int
	n := len(arr)

	for i := 0; i < n; i++ {
		pos = i

		for j := i + 1; j < n; j++ {
			if arr[j] < arr[pos] {
				pos = j
			}
		}
		arr[i], arr[pos] = arr[pos], arr[i]
	}
}

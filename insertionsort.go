package main

import "fmt"

func main() {
	arr := []int{6, 15, 9, 2, 3}
	fmt.Println(arr)
	insertionsort(arr)
	fmt.Println(arr)
}

func insertionsort(arr []int) {
	n := len(arr)
	var cur, pos int

	for i := 0; i < n; i++ {
		cur = arr[i]
		pos = i

		for pos > 0 && arr[pos-1] > cur {
			arr[pos] = arr[pos-1]
			pos--
		}
		arr[pos] = cur
	}
}

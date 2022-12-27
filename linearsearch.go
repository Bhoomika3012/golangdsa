package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50}
	fmt.Println(linearsearch(arr, 40))
}
func linearsearch(arr []int, key int) int {

	for i, v := range arr {
		if key == v {
			return i
		}

	}
	return -1
}

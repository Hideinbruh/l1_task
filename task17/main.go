package main

import "fmt"

func binarySearch(arr []int, target int) int {
	low := 0
	high := len(arr) - 1

	for low <= high {
		mid := (low + high) / 2

		if arr[mid] < target {
			low = mid + 1
		} else if arr[mid] > target {
			high = mid - 1
		} else {
			return mid
		}
	}

	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9}
	target := 5

	index := binarySearch(arr, target)

	if index == -1 {
		fmt.Printf("%d не найден в массиве %v\n", target, arr)
	} else {
		fmt.Printf("%d найден в массиве %v на позиции %d\n", target, arr, index)
	}
}

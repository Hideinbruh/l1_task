package main

import "fmt"

// Удалить i-ый элемент из слайса.

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	i := 2
	numbers = append(numbers[:i], numbers[i+1:]...)
	fmt.Println(numbers)
}

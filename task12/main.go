package main

import (
	"fmt"
)

// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее
//собственное множество

func main() {
	words := []string{"cat", "cat", "dog", "cat", "tree"}
	set := make(map[string]bool)

	for _, word := range words {
		set[word] = true
	}
	for word := range set {
		fmt.Println(word)
	}
}

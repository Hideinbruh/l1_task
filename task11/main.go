package main

import "fmt"

// Реализовать пересечение двух неупорядоченных множеств.

func main() {
	firstSet := []int{5, 52, 35, 33, 2, 0, 11, 153, 123, 6, 87}
	secondSet := []int{37, 89, 90, 2, 3, 15, 5, 23, 87, 68, 27}
	intersection := make([]int, 0)

	for _, elemFirstSet := range firstSet {
		for _, elemSecondSet := range secondSet {
			if elemFirstSet == elemSecondSet {
				intersection = append(intersection, elemFirstSet)
			} else {
				continue
			}
		}
	}
	fmt.Println(intersection)
}

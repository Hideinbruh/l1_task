package main

import "fmt"

//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0,
//15.5, 24.5, -21.0, 32.5. Объединить данные значения в группы с шагом в 10
//градусов. Последовательность в подмножноствах не важна.

func main() {
	temps := []float64{-25, 4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	groups := make(map[int][]float64)

	for _, temp := range temps {
		key := int(temp/10) * 10
		groups[key] = append(groups[key], temp)
	}

	for key, group := range groups {
		fmt.Printf("%d: %v\n", key, group)
	}

}

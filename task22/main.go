package main

import (
	"fmt"
	"math"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает две
//числовых переменных a,b, значение которых > 2^20.

func main() {
	a := math.Pow(2, 21)
	b := math.Pow(2, 22)

	m := a * b
	fmt.Println("a * b =", m)

	m = a / b
	fmt.Println("a / b =", m)

	m = a + b
	fmt.Println("a + b = =", m)

	m = a - b
	fmt.Println("a - b =", m)
}

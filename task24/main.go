package main

import "fmt"

//Разработать программу нахождения расстояния между двумя точками, которые
//представлены в виде структуры Point с инкапсулированными параметрами x,y и
//конструктором.

type Point struct {
	a float64
	b float64
}

func NewPoint(a, b float64) *Point {
	return &Point{
		a: a,
		b: b,
	}
}

func main() {
	points := NewPoint(2.0, 5.0)
	var distance float64
	if points.a < points.b {
		distance = points.b - points.a
		fmt.Println(distance)
	} else if points.a > points.b {
		distance = points.a - points.b
		fmt.Println(distance)
	} else {
		fmt.Println("Distance is zero")
	}
}

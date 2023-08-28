package main

import "fmt"

//Разработать программу, которая в рантайме способна определить тип
//переменной: int, string, bool, channel из переменной типа interface{}.

func main() {
	a := "Hello"
	TypeDef(a)
}

func TypeDef(a interface{}) {
	switch a.(type) {
	case int:
		fmt.Printf("type of variable is int")
	case string:
		fmt.Printf("type of variable is string")
	case bool:
		fmt.Printf("type of variable is bool")
	case chan int:
		fmt.Printf("type of variable is channel int")
	default:
		fmt.Println("unknown type")
	}
}

package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры
//Human (аналог наследования).

type Human struct {
	Name string
	Age  int
}

func (h *Human) SayHello() {
	fmt.Printf("Hello, my name is %s, i'm %d years old\n", h.Name, h.Age)
}

// Action если передать структуру Human с именем, то обращаться к полям Human можно через a.Human.Name
// если передать структуру Human без имени, то обращаться к полям Human можно напрямую a.Name
type Action struct {
	// передаем указатель на структуру, т.к. в этом случае изменения структуры будут видны за пределами функции, в которой поля изменяются
	// при передаче указателя расходуется меньше памяти, т.к. передается не копия, а адрес, в которой она лежит
	*Human
}

func (a *Action) Run() {
	fmt.Printf("%s is running\n", a.Name)
}

func main() {
	human := &Human{
		Name: "John",
		Age:  25,
	}
	action := &Action{Human: human}
	action.SayHello()
	action.Run()

}

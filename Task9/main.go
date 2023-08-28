package main

import (
	"fmt"
	"sync"
	"time"
)

// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из
//массива, во второй — результат операции x*2, после чего данные из второго
//канала должны выводиться в stdout.

func main() {
	num := [5]int{1, 2, 3, 4, 5}
	ch1 := make(chan int, len(num))
	ch2 := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for _, elem := range num {
			ch1 <- elem
			time.Sleep(time.Second)
		}
		close(ch1)
		defer wg.Done()
	}()

	wg.Add(1)
	go func() {
		for elem := range ch1 {
			ch2 <- elem * 2
		}
		close(ch2)
		wg.Done()
	}()

	for res := range ch2 {
		fmt.Println(res)
	}
	fmt.Println("App stopped")

}

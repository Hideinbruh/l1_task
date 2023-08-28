package main

import (
	"fmt"
	"sync"
	"time"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел
//взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	square := []int{}
	var wg sync.WaitGroup // создаем wait группу, чтобы программа не завершилась до исполнения всех горутин
	var mu sync.Mutex     // создаем mutex для защиты доступа к слайсу

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, num := range numbers {
			multiply := num * num
			mu.Lock() // защищаем доступ к square
			square = append(square, multiply)
			mu.Unlock()
		}
	}()
	wg.Add(1)
	go func() {
		<-time.After(time.Second) // даем первой горутине успеть внести изменения в square
		defer wg.Done()
		mu.Lock()
		for _, sq := range square {
			fmt.Println(sq)
		}
		mu.Unlock()
	}()
	wg.Wait()

	// второй вариант решения через каналы
	OtherVar()
}

func OtherVar() {
	numbers := []int{2, 4, 6, 8, 10}
	square := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, num := range numbers {
			square <- num * num
			time.Sleep(time.Second)
		}
		close(square)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		for sq := range square {
			fmt.Println(sq)
		}
	}()
	wg.Wait()
}

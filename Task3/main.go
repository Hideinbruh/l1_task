package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел
//взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	results := make(chan int, len(numbers)) // канал, в который будем прокидывать квадраты значений, взятых из numbers
	var wg sync.WaitGroup                   // waitgroup для ожидания завершения работы всех горутин

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			results <- n * n
		}(num)
	}
	wg.Wait()
	close(results)
	sum := 0
	for result := range results {
		sum += result
	}
	fmt.Printf("Сумма квадратов %v равна %v\n", numbers, sum)
	// другой вариант решения
	Var()

}

func Var() {
	numbers := []int{2, 4, 6, 8, 10}
	var result []int
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			multiply := n * n
			mu.Lock()
			result = append(result, multiply)
			mu.Unlock()
		}(num)
	}
	wg.Wait()
	sum := 0
	for _, results := range result {
		sum += results
	}
	fmt.Printf("Сумма квадратов %v равна %v\n", numbers, sum)
}

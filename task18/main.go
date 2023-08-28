package main

import (
	"fmt"
	"sync"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в
//конкурентной среде. По завершению программа должна выводить итоговое
//значение счетчика.

type Counter struct {
	count int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func main() {
	coun := &Counter{count: 0}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			coun.Increment()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(coun.count)
}

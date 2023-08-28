package main

import (
	"fmt"
	"sync"
	"time"
)

// Реализовать конкурентную запись данных в map.

func main() {
	// именно конкурентную запись реализовать невозможно, так как map не является потокобезопасной и при одновременной
	// записи один и тот же ключ может быть перезаписан несколькими горутинами, что приведет к потере данных.
	// По этой причине записывать данные горутинами в map можно, используя мьютексы для ограничения доступа к map
	var mu sync.Mutex
	var wg sync.WaitGroup
	data := make(map[int]interface{})
	numWorkers := 3
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go Writer(i, data, &mu, &wg)
	}
	wg.Wait()
	fmt.Println("App stopped")
}

func Writer(workerId int, data map[int]interface{}, mu *sync.Mutex, wg *sync.WaitGroup) {
	mu.Lock()
	for i := 0; i < 5; i++ {
		value := i * i
		data[i] = value
		fmt.Printf("Worker %d sent into map: key %d and value %d\n", workerId, i, value)
		time.Sleep(time.Second)
	}
	fmt.Printf("Worker %d stopped it's work\n", workerId)
	wg.Done()
	mu.Unlock()
}

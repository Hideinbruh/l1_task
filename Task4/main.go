package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток). Реализовать
//набор из N воркеров, которые читают произвольные данные из канала и
//выводят в stdout. Необходима возможность выбора количества воркеров при
//старте.
//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать
//способ завершения работы всех воркеров.

func main() {

	ch := make(chan string)
	numWorkers := os.Args[1]
	nw, err := strconv.Atoi(numWorkers)
	var wg sync.WaitGroup

	if err != nil {
		log.Fatalf("Invalid number workers", err)
		return
	}
	go Writer(ch)

	for i := 0; i < nw; i++ {
		wg.Add(1)
		go Worker(i, ch, &wg)
	}
	wg.Wait()
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	fmt.Println("Все воркеры завершили работу")
	close(ch)

}

func Writer(ch chan string) {
	for {
		data := generateData()
		ch <- data
		fmt.Println(ch)
		time.Sleep(time.Second)
	}
}
func generateData() string {
	LRunes := []rune("ABCDEFG")
	b := make([]rune, len(LRunes))
	for i := range b {
		b[i] = LRunes[rand.Intn(len(LRunes))]
	}
	return string(b)
}

func Worker(id int, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for d := range ch {
		fmt.Printf("%d worker recived data from %T: %s\n", id, ch, d)
	}
	fmt.Printf("%d worker stopped", id)
}

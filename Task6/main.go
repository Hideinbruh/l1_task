package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Реализовать все возможные способы остановки выполнения горутины

//1. Передать канал в горутину и закрыть его спустя какое-то время
//2.Использование context
//3.Использование context с методом cancel()
//4.остановить main

func main() {
	//1.Использование канала
	ch := make(chan interface{})
	go func(ch chan interface{}) {
		for i := 0; ; i++ {
			select {
			default:
				fmt.Println("Sent:", i)
				time.Sleep(time.Second)
			case <-ch:
				fmt.Println("Goroutine stopped")
				return
			}
		}
	}(ch)
	time.Sleep(time.Second * 5) // ждем 5 секунд и отправляем сообщение в канал для закрытия горутины
	close(ch)
	time.Sleep(time.Second) // ждем остановки горутины
	fmt.Println("first way completed")

	//2.использование context с временем
	ctx, _ := context.WithTimeout(context.Background(), time.Second*5)
	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup, ctx context.Context) {
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine stoppped")
				wg.Done()
				return
			default:
				fmt.Println("Sent: ", i)
				time.Sleep(time.Second)
			}
		}
	}(&wg, ctx)
	wg.Wait()
	fmt.Println("second way completed")

	//3. Использование context с методом cancel()
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Println("Goroutine stopped")
				return
			default:
				fmt.Println("Sent:", i)
				time.Sleep(time.Second)
			}
		}
	}(ctx)
	<-time.After(time.Second * 5)
	cancel()
	time.Sleep(time.Second)
	fmt.Println("third way completed")
}

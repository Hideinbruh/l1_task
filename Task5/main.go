package main

import (
	"context"
	"fmt"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в
//канал, а с другой стороны канала — читать. По истечению N секунд программа
//должна завершаться

func sendData(ch chan int, ctx context.Context) {
	for i := 0; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Println("Stop sending data")
			return
		case ch <- i:
			fmt.Println("Sent:", i)
			time.Sleep(time.Second)
		}
	}
}

func main() {
	ch := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	go sendData(ch, ctx)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Time is up")
			return
		case data := <-ch:
			fmt.Println("Received:", data)

		}
	}
}

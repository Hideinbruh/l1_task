package main

import (
	"fmt"
	"time"
)

func sleep(t time.Duration) {
	time.Sleep(time.Second * t)
}

func main() {
	fmt.Println("App starts")
	fmt.Println("Waiting for 5 seconds")
	sleep(5)
	fmt.Println("App stopped")
}

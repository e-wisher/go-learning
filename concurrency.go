package main

import (
	"fmt"
	"sync"
	"time"
)

var waitgroup sync.WaitGroup

func displayNumber(routine string) {
	for i := 0; i < 10; i++ {
		fmt.Println(routine, ":", i)
		time.Sleep(time.Second)
	}
	waitgroup.Done()
}

func ping(pings chan<- string, msg string) {
	pings <- msg
	time.Sleep(time.Second)
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {

	waitgroup.Add(1)
	go displayNumber("Goroutine")
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "Ping pong!")
	pong(pings, pongs)
	fmt.Println(<-pongs)
	waitgroup.Wait()
}

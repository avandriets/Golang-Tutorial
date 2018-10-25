package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

// Двунаправленный канал c chan string
func ponger(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

// Однонаправленный канал c <- chan string
func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// Синхронный канал make(chan string) !!!
	var c chan string = make(chan string)

	go printer(c)
	go ponger(c)
	go pinger(c)

	var input string
	fmt.Scanln(&input)
}

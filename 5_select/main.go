package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan int)

	go func1(ch1)
	go func2(ch2)

	for {
		select {
		case msg := <-ch1:
			fmt.Printf("Channel1: %s\n", msg)
		case msg := <-ch2:
			fmt.Printf("Channel2: %d\n", msg)
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout: no data received within 3 seconds")
		}
	}

}

func func1(ch chan string) {
	ch <- "hello"
}

func func2(ch chan int) {
	ch <- 1
}

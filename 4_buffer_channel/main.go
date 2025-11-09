package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	go producer(ch)
	consumer(ch)
}

func producer(ch chan<- int) {
	defer close(ch)

	for i := range 10 {
		fmt.Printf("%d 生成\n", i)
		ch <- i
	}
}

func consumer(ch <-chan int) {
	for i := range ch {
		fmt.Printf("%d 消費\n", i)
	}
}

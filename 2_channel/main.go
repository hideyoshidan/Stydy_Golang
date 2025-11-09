package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	// sender
	go func() {
		for i := range 10 {
			fmt.Printf("%d: sending... \n", i)
			ch <- i
		}
		close(ch)
	}()

	// reciver
	for n := range ch {
		fmt.Printf("%d recieved! \n", n)
	}
}

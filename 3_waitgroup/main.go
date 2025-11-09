package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	for i := range 5 {
		wg.Add(1) // Add(n)のnはgoroutineの数

		go func(id int) {
			defer wg.Done()

			fmt.Printf("%d: start... \n", i)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("%d: finished! \n", i)
		}(i)
	}

	wg.Wait()
	fmt.Print("all Task is done.")
}

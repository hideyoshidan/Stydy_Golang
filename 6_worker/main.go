package main

import (
	"fmt"
	"sync"
)

/*
*
* 必ず、 worker起動 -> channelに送信の順にしないと危ない。
もしくは、buffer size >= task数の場合はok
*/
func main() {
	maxWorker := 3
	taskNum := 100

	var wg sync.WaitGroup
	taskCh := make(chan int, 10)

	// workerを3つ起動
	for i := range maxWorker {
		wg.Add(1)
		go work(i, &wg, taskCh)
	}

	// channelに100個送信
	send(taskCh, taskNum)

	wg.Wait()
}

func work(id int, wg *sync.WaitGroup, ch <-chan int) {
	defer wg.Done()
	for t := range ch {
		fmt.Printf("%d 個目のtask on worker:%d\n", t, id)
	}
}

func send(ch chan int, taskNum int) {
	for i := range taskNum {
		ch <- i
	}
	close(ch)
}

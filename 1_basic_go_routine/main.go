package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		go printMsg(i, fmt.Sprintf("%d番目のmessageです。", i))
	}
	time.Sleep(1 * time.Second)

}

func printMsg(i int, msg string) {
	fmt.Printf("%d : %s\n", i, msg)
}

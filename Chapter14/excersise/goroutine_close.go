package main

import "fmt"

func goroutineClose() {
	ch := make(chan int)
	go telb(ch)
	for {
		c, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(c)
	}
}

func telb(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

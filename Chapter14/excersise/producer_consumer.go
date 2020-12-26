package main

import "fmt"

func producerConsumer() {
	ch := make(chan int)
	done := make(chan bool)
	go func(ch chan int) {
		for i := 0; i < 100; i = i + 10 {
			ch <- i
		}
		close(ch)
	}(ch)
	// <-chan 表示只能从信道里发出数据，对程序来说是只读
	// chan<- 表示信道只能从外面接收数据，对程序来说是只写
	go func(ch <-chan int, done chan<- bool) {
		for c := range ch {
			fmt.Println(c)
		}
		done <- true
	}(ch, done)
	<-done
}

package main

import (
	"fmt"
	"time"
)

func goroutineSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go pump1(ch1)
	go pump2(ch2)
	go suck0(ch1, ch2)
	time.Sleep(10 * time.Millisecond)
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

// select 语句实现了一种监听模式，通常用在（无限）循环中；在某种情况下，通过 break 语句使循环退出
func suck0(ch1 chan int, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on ch1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on ch2: %d\n", v)
		}
	}
}

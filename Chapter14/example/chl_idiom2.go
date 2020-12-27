package main

import (
	"fmt"
	"time"
)

func chlIdiom2() {
	ch := pump()
	suck(ch)
	fmt.Println("test2")
	time.Sleep(10)
	fmt.Println("test3")
}

func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()
	fmt.Println("test1")
	return ch
}

func suck(ch chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
}

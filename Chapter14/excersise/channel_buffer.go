package main

import (
	"fmt"
	"time"
)

func chlBuffer() {
	ch := make(chan int, 1)
	go func() {
		time.Sleep(15 * time.Second)
		x := <-ch
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	ch <- 10
	fmt.Println("sent", 10)
}

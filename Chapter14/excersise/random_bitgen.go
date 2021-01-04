package main

import (
	"fmt"
	"time"
)

func randomBitGen() {
	c := make(chan int)
	go func() {
		for {
			fmt.Print(<-c, " ")
		}
	}()
	go func() {
		for {
			select {
			case c <- 0:
			case c <- 1:
			}
		}
	}()
	time.Sleep(time.Second)
}

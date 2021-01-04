package main

import "fmt"

func fibonacciSelect(c chan int, quit chan bool) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func goFibonacciSelect() {
	ch := make(chan int)
	quit := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-ch)
		}
		quit <- true
	}()
	fibonacciSelect(ch, quit)
}

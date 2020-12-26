package main

import "fmt"

func gosum() {
	a := 2
	b := 3
	ch := make(chan int)
	go func(a, b int, ch chan int) {
		ch <- a + b
	}(a, b, ch)
	fmt.Printf("%d + %d = %d", a, b, <-ch)
}

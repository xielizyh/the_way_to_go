package main

import "fmt"

func goroutinePanic() {
	ch := make(chan int)
	go tela(ch)
	for {
		fmt.Println(<-ch)
	}
}

func tela(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

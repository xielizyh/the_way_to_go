package main

import "fmt"

func fibonacci(n int) int {
	if n <= 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibnterms(term int, c chan int) {
	for i := 0; i < term; i++ {
		c <- fibonacci(i)
	}
	close(c)
}

func goFibonacci() {
	term := 25
	ch := make(chan int)
	go fibnterms(term, ch)
	for v := range ch {
		fmt.Println(v)
	}
}

/*
func goFibonacci() {
	ch := make(chan int)
	go fibonacci(5, ch)
	fmt.Println(<-ch)
}

func fibonacci(n int, ch chan int) {
	c := make(chan int)
	var res int
	if n <= 1 {
		ch <- 1
		return
	}
	go fibonacci(n-1, c)
	res = <-c
	go fibonacci(n-2, c)
	res += <-c
	ch <- res
}
*/

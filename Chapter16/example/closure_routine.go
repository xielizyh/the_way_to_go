package main

import (
	"fmt"
	"time"
)

func closureRoutine() {
	values := []int{10, 11, 12, 13, 14}

	// A version
	for ix := range values {
		func() {
			fmt.Printf("%d ", ix)
		}()
	}
	fmt.Println()
	// B version 协程可能在循环结束后还没有开始执行，所以ix的值为4
	// https://studygolang.com/articles/21721
	for ix := range values {
		go func() {
			fmt.Printf("%d ", ix)
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println()
	// C version
	for ix := range values {
		go func(ix interface{}) {
			fmt.Printf("%d ", ix)
		}(ix)
	}
	time.Sleep(1 * time.Second)
	fmt.Println()
	// D version
	for _, val := range values {
		go func(val interface{}) {
			fmt.Printf("%d ", val)
		}(val)
	}
	time.Sleep(1 * time.Second)
}

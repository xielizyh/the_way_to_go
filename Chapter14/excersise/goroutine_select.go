package main

import "fmt"

func goroutineSelect() {
	var exit = false
	ch := make(chan int)
	done := make(chan bool)
	go telc(ch, done)
	for !exit {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-done:
			exit = true
			fmt.Println("exit")
		}
	}
}

func telc(ch chan int, done chan bool) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	done <- true
}

/* 以下代码会导致死锁
https://www.cnblogs.com/faithfu/p/12067753.html
func goroutineSelect() {
	ch := make(chan int)
	done := make(chan bool)
	go telc(ch, done)
	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-done:
			fmt.Println("exit")
			// for和select一同使用，有坑
			// break只能跳出select，无法跳出for
			// 然后继续下一次select，而此时telc线程已经执行完成，
			// 故会导致死锁（所有线程均睡眠了）
			break
		}
		fmt.Println("here")
	}
}

func telc(ch chan int, done chan bool) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	done <- true
}
*/

/*
// 协程调度的本质是IO阻塞，即IO阻塞时才会发生调度
func goroutineSelect() {
	ch := make(chan int)
	done := make(chan bool)
	go telc(ch, done)
	fmt.Println("here2")
	go func() {
		fmt.Println("here3")
		for {
			fmt.Println(<-ch)
		}
	}()
	fmt.Println("here0")
	<-done
}

func telc(ch chan int, done chan bool) {
	fmt.Println("here4")
	for i := 0; i < 10; i++ {
		ch <- i
	}
	fmt.Println("here1")
	done <- true
}
*/

/*
// 注意以下程序会导致死锁
func goroutineSelect() {
	ch := make(chan int)
	done := make(chan bool)
	go telc(ch, done)
	fmt.Println("here2")
	go func() {
		fmt.Println("here3")
		// 这里没有循环打印，会导致死锁
		fmt.Println(<-ch)
	}()
	fmt.Println("here0")
	<-done
}

func telc(ch chan int, done chan bool) {
	fmt.Println("here4")
	for i := 0; i < 10; i++ {
		ch <- i
	}
	fmt.Println("here1")
	done <- true
}
*/

package main

import (
	"fmt"
	"time"
)

// 协程切换的本质是异步io函数(busy活着blocking)

// chlBlock 可以得到输出“received 10”
// 因为通道发送先准备好，故执行“x := <-ch”不会阻塞，
// 输出“received 10”后，协程执行完毕，返回发送协程继续执行
func chlBlock() {
	// 无缓冲通道，发送操作，在接收者准备好之前是阻塞的
	ch := make(chan int)
	go func() {
		time.Sleep(5 * time.Second)
		x := <-ch
		// time.Sleep(1)
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	ch <- 10
	fmt.Println("sent", 10)
}

// chlBlock1 不会得到输出“received 10”
// 因为通道接收先准备好，故执行“ch <- 10”不会阻塞，
// 程序直接结束，导致通道接收协程不能执行
func chlBlock1() {
	// 无缓冲通道，发送操作，在接收者准备好之前是阻塞的
	ch := make(chan int)
	go func() {
		// time.Sleep(5 * time.Second)
		x := <-ch
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	ch <- 10
	fmt.Println("sent", 10)
}

// chlBlock2 错误的方式，造成死锁
func chlBlock2() {
	// 无缓冲通道，发送操作，在接收者准备好之前是阻塞的
	ch := make(chan int)
	fmt.Println("sending", 10)
	ch <- 10 // 死锁
	fmt.Println("sent", 10)
	go func() {
		time.Sleep(5 * time.Second)
		x := <-ch
		fmt.Println("received", x)
	}()
}

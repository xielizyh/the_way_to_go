package main

import "fmt"

// https://zhuanlan.zhihu.com/p/74047342
// https://studygolang.com/articles/9592
// https://studygolang.com/articles/4474
// go的协程是基于线程的，是应用层模拟的线程
// 线程是操作系统内核层，协程是应用层实现的
// go协程维护了一组数据结构和n个线程，真正的执行还是线程，
// 协程执行的代码被扔进一个待执行队列中，有这n个线程从队列中拉出来执行。
// 并行是一种通过使用多处理器（多核）以提高速度的能力
// 并发可以在一个处理器或者内核上使用多个线程来执行任务，
// 但是只有同一个程序在某个时间点同时运行在多核或者多处理器上才是真正的并行。
// 并发程序可以是并行的，也可以不是
// “并发关乎结构，并行关乎执行”
func main() {
	fmt.Println("------Example 14.6-----")
	// chlIdiom2()
	fmt.Println("------Example 14.10-----")
	goroutineSelect()
}

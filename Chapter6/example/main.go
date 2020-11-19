package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func main() {
	// 6.9
	fmt.Println("------Example 6.9------")
	var f = Adder()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
	fmt.Println()
	// 6.10
	fmt.Println("------Example 6.10------")
	where := func() {
		// skip是要提升的堆栈帧数，0-当前函数，1-上一层函数，....
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
	log.SetFlags(log.Llongfile)
	log.Print("")
	where1 := log.Print
	where1()
	// 6.11
	fmt.Println("------Example 6.11------")
	start := time.Now()
	for i := 0; i < 10000000; i++ {
		f(4000)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("took this amount of time: %s\n", delta)
	// 6.12
	fmt.Println("------Example 6.12------")
	var result uint64 = 0
	start = time.Now()
	for i := 0; i < LIM; i++ {
		result = Fibonacci(i)
		fmt.Printf("Fibonacci(%d): %d\n", i, result)
	}
	end = time.Now()
	delta = end.Sub(start)
	fmt.Printf("took this amount of time: %s\n", delta)
}

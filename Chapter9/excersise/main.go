package main

import (
	"container/list"
	"fmt"
	"unsafe"

	"excersise/even"
	"excersise/fibo"
	"excersise/greetings"
)

func main() {
	// 9.1
	fmt.Println("------Excersise 9.1------")
	list1 := list.New()
	list1.PushBack(101)
	list1.PushBack(202)
	list1.PushBack(303)
	for e := list1.Front(); e != nil; e = e.Next() {
		fmt.Printf("%d ", e.Value)
	}
	fmt.Println()
	// 9.2
	fmt.Println("------Excersise 9.2------")
	var tmp2 int = 12
	fmt.Printf("unsafe.sizeof(int)=%d\n", unsafe.Sizeof(tmp2))
	// 切片源码定义
	// type slice struct {
	// 	array unsafe.Pointer // 元素指针
	// 	len   int            // 长度
	// 	cap   int            // 容量
	// }
	sli2 := make([]int, 9, 20)
	// len: &s => pointer => uintptr => pointer => *int => int
	len := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&sli2)) + uintptr(8)))
	// cap: &s => pointer => uintptr => pointer => *int => int
	cap := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&sli2)) + uintptr(16)))
	fmt.Println("len=", len, "cap=", cap)

	// 9.3
	fmt.Println("------Excersise 9.3------")
	// 一定要注意go.mod里的module名字，必须为文件夹名即excersise
	name := "David"
	fmt.Println(greetings.GoodDay(name))
	fmt.Println(greetings.GoodNight(name))
	if greetings.IsAm() {
		fmt.Println("Good morning", name)
	} else if greetings.IsAfternoon() {
		fmt.Println("Good afternoon", name)
	} else if greetings.IsEvening() {
		fmt.Println("Good evening", name)
	} else {
		fmt.Println("Good night", name)
	}
	// 9.4
	fmt.Println("------Excersise 9.4------")
	for i := 0; i < 10; i++ {
		if even.IsEven(i) {
			fmt.Println(i, "is even")
		} else {
			fmt.Println(i, "is odd")
		}
	}
	// 9.5
	fmt.Println("------Excersise 9.5------")
	for i := 0; i <= 10; i++ {
		fmt.Printf("fibonacci(%d) is: %d\n", i, fibo.Fibonacci(i))
	}

}

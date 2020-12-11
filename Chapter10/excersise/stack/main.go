package main

import (
	"fmt"
	stackarr "stack/stack_arr"
	stackstru "stack/stack_struct"
)

func main() {
	stackArr := stackarr.Stack{0}
	(&stackArr).Push(11)
	(&stackArr).Push(12)
	fmt.Println("After push, statck is", stackArr)
	// data := (&stackArr).Pop()
	// 如下，我们并没有在指针上进行调用，go为我们做了探测工作，
	data := stackArr.Pop()
	fmt.Println("The pop data is", data)
	fmt.Println("After pop, stack is", stackArr)

	// len函数对于指针来说也是求得其指向变量的长度
	// tmp := [3]int{11, 12}
	// ptr := &tmp
	// fmt.Println(len(tmp))
	// fmt.Println(len(ptr))
	// fmt.Println(len(*ptr))

	fmt.Println("--------------------------------")
	stackStru := new(stackstru.Stack)
	stackStru.Push(55)
	stackStru.Push(66)
	fmt.Println("After push, statck is", stackStru)
	data = stackStru.Pop()
	fmt.Println("The pop data is", data)
	fmt.Println("After pop, stack is", stackStru)
}

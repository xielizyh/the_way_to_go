package main

import "fmt"

func main() {
	// 7.1
	fmt.Println("------Excersise 7.1------")
	var arr1 [3]int
	ArrayValue(arr1)
	for i := range arr1 {
		fmt.Printf("%d ", arr1[i])
	}
	fmt.Println()
	// 7.2
	fmt.Println("------Excersise 7.2------")
	var arr2 [50]int
	FibonacciArray(&arr2)
	for i := range arr2 {
		fmt.Printf("%d ", arr2[i])
	}
	fmt.Println()
	// 7.3
	fmt.Println("------Excersise 7.3------")
	value := FibonacciFuncArray(10)
	for i := range value {
		fmt.Printf("%d ", value[i])
	}
	fmt.Println()
	// 7.5
	fmt.Println("------Excersise 7.5------")
	var array [10]byte = [10]byte{'A', 'B', 'C', 'D', 'E', 'F'}
	sl := make([]byte, 8, 16)
	for i := 0; i < len(sl); i++ {
		sl[i] = 'a' + byte(i)
	}
	sl = Append(sl, array[:])
	fmt.Println(sl)

	// append 超过slice底层数组的容量问题
	// 每个slice在底层都有一个数组作为支撑
	// 切片可以理解为“动态数组”，即数组长度可以改变
	// https://studygolang.com/articles/29977?fr=sidebar
	// 1. 容量够
	slice3 := make([]int, 2)
	slice4 := append(slice3, 100)
	fmt.Println("slice3", slice3)
	fmt.Println("slice4", slice4)
	fmt.Printf("slice3 address = %p\n", &slice3)
	fmt.Printf("slice4 address = %p\n", &slice4)
	// 2. 容量不够
	slice5 := make([]int, 1)
	slice6 := append(slice5, 100)
	fmt.Println("slice5", slice5)
	fmt.Println("slice6", slice6)
	fmt.Printf("slice5 address = %p\n", &slice5)
	fmt.Printf("slice6 address = %p\n", &slice6)

	slice7 := make([]int, 1)
	slice7[0] = 2
	slice8 := make([]int, 1)
	slice8 = slice7
	fmt.Println(slice8)

	// start := time.Now()
	// tmp := 0
	// for {
	// 	tmp++
	// 	if tmp > 1e8 {
	// 		break
	// 	}
	// }
	// end := time.Now()
	// delta := end.Sub(start)
	// fmt.Printf("took this amount of time: %s\n", delta)
}

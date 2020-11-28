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

	// 7.6
	fmt.Println("------Excersise 7.6------")
	var buf []byte = []byte{11, 22, 33, 44, 55}
	buf1, buf2 := buf[:2], buf[2:]
	fmt.Println(buf1, buf2)

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
	// 7.7
	fmt.Println("------Excersise 7.7------")
	arrF := []float32{1.1, 2.2, 3.3, 4.4, 5.5, 6.6}
	fmt.Println(SumArray(arrF))
	// fmt.Println(SumArray1(arrF[:]...))
	// 传入参数要进行展开
	fmt.Println(SumArray1(arrF...))
	fmt.Println(SumAndAverage(100, 77.234))

	// 7.8
	fmt.Println("------Excersise 7.8------")
	fmt.Println(MinSlice([]int{0, 1, 2, 3, 4, 5, 6, 7}...))
	fmt.Println(MaxSlice([]int{0, 1, 2, 3, 4, 5, 6, 7}...))
	fmt.Println(MinSlice(0, 1, 2, 3, 4, 5, 6, 7))
	fmt.Println(MaxSlice(0, 1, 2, 3, 4, 5, 6, 7))

	// 7.9
	fmt.Println("------Excersise 7.9------")
	slice10 := make([]int, 4)
	slice9 := func(slice []int, factor int) []int {
		newSlice := make([]int, len(slice)*factor)
		copy(newSlice, slice)
		slice = newSlice
		return slice
	}(slice10, 2)
	fmt.Println(len(slice10), len(slice9))

	// 7.10
	fmt.Println("------Excersise 7.10------")
	// 注意数组和切片的区别，切片不需要说明长度，数组声明就需要定义长度
	slice00 := []int{1, 2, 3}
	// go中数组和C/C++不一样，它是一种值类型，即类似int这种基本类型
	array00 := [...]int{1, 2, 3}
	fmt.Printf("The type of slice00, array00 is %T, %T\n", slice00, array00)
	// arr11类型是*[2]int，arr12类型是[2]int
	var arr11 = new([2]int)
	arr12 := *arr11
	arr12[0] = 100
	fmt.Println(*arr11, arr12)
	var arr14 = new([]int)
	var arr15 = make([]int, 0)
	arr15 = *arr14
	fmt.Printf("The type of arr14, arr15 is %T, %T\n", arr14, arr15)

	slice11 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice11 = Filter(slice11, IsEven)
	fmt.Println(slice11)
}

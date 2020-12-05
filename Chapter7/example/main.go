package main

import (
	"fmt"
)

func main() {
	// 7.7
	fmt.Println("------Example 7.7------")
	var arr1 [6]int
	var slice1 []int = arr1[2:5]

	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	// 切片只能向后移动，故slice1容量是4不是6
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// panic: runtime error: slice bound out of range
	// slice1 = slice1[0:7]

	// 7.8
	fmt.Println("------Example 7.8------")
	var slice2 []int = make([]int, 10, 20)
	for i := 0; i < len(slice2); i++ {
		slice2[i] = 5 * i
	}
	for i := 0; i < len(slice2); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice2[i])
	}
	slice21 := slice2[:20]
	for i := 0; i < cap(slice21); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice21[i])
	}
	fmt.Printf("The length of slice2 is %d\n", len(slice2))
	fmt.Printf("The capacity of slice2 is %d\n", cap(slice2))

	// new 只分配内存
	// make 只能用于 slice、map 和 channel 的初始化
	// https: //www.cnblogs.com/lurenq/p/12013250.html
	// make 只能用来分配及初始化类型为 slice、map、chan 的数据。new 可以分配任意类型的数据；
	// new 分配返回的是指针，即类型 *Type。make 返回引用，即 Type；
	// new 分配的空间被清零。make 分配空间后，会进行初始化；
	type Student struct {
		name string
		age  int
	}
	var s *Student
	// 类似C语言中的malloc函数
	s = new(Student)
	s.name = "xieli"
	fmt.Println(s)

	type test struct {
		// 空数组
		tmp []int
	}
	test1 := test{[]int{1, 2, 3}}
	fmt.Println(test1.tmp)

	var test2 []int
	test2 = []int{3, 4, 5}
	fmt.Println(test2)

	// golang 当变量(包括结构体中的变量)或函数的首字母大写的时候，变量/函数会被从包中导出（在包外部可见，或者说公有的）
	// 7.10
	fmt.Println("------Example 7.10------")
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for _, s := range seasons {
		fmt.Printf("%s\n", s)
	}
	for ix := range seasons {
		fmt.Printf("%d\n", ix)
	}
	items := []int{1, 2, 3, 4, 5, 6, 7}
	for _, item := range items {
		// item 只是 items 某个索引位置的值的一个拷贝，不能用来修改 items 该索引位置的值
		item *= 2
	}
	fmt.Println(items)

	for idx := range items {
		items[idx] *= 2
	}
	fmt.Println(items)

	sum := func(num ...int) int {
		sum := 0
		for _, value := range num {
			sum += value
		}
		return sum
	}(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Printf("sum = %d\n", sum)

	// 切片可以反复扩展直到占据整个相关数组
	var ar = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	a := ar[2:2]
	fmt.Println("The length of a is", len(a))
	fmt.Println("The capcity of a is", cap(a))
	a = ar[2:3]
	fmt.Println("The length of a is", len(a))
	fmt.Println("The capcity of a is", cap(a))

	// 7.12
	fmt.Println("------Example 7.12------")
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)
	// 拷贝个数是 slFrom 和 dslTo 的长度最小值
	n := copy(slTo, slFrom)
	fmt.Println(slTo)
	fmt.Printf("Copied %d elements\n", n)

	sl3 := []int{1, 2, 3}
	// 如果 s 的容量不足以存储新增元素，
	// append 会分配新的切片来保证已有切片元素和新增元素的存储
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)

	sl4 := make([]byte, 2)
	sl5 := make([]byte, 4)
	//  append 操作时，一定要将 append 的结果重新赋值给一个 slice，防止由于底层数组的变更导致的数据错误
	sl6 := AppendByte(sl4, sl5...)
	fmt.Println(sl4, sl6)
	// https://studygolang.com/articles/13405?fr=sidebar
	// 7.13
	fmt.Println("------Example 7.13------")
	s17 := "hello world"
	for i, c := range s17 {
		fmt.Printf("%d:%c ", i, c)
	}
	// 字符串本质上是一个字节数组
	c := []byte(s17)
	fmt.Println(c)
	c = append(c, " string test"...)
	fmt.Println(string(c))
	// \u 表示Unicode码（码位），utf-8表示编码规则（将「码位」转换为字节序列的规则）
	s18 := "\u00ff\u754c"
	for i, c := range s18 {
		fmt.Printf("%d:%c ", i, c)
	}
	// Go 语言中的字符串是不可变的，报错cannot assign to s17[0]
	// s17[0] = 'c'
	// 必须先将字符串转换成字节数组
	c[0] = 'c'
	s19 := string(c)
	fmt.Println()
	fmt.Println(s19)
	// 切片的底层指向一个数组，只有在没有任何切片指向的时候，底层的数组内存才会被释放
	// 想要避免这个问题，可以通过拷贝我们需要的部分到一个新的切片
}

// AppendByte 追加字节
func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	// 原始切片容量不够，重新分配内存
	if n > cap(slice) {
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

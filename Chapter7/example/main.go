package main

import "fmt"

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

}

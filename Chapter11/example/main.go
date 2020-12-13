package main

import (
	"fmt"

	mysort "example/mysort"
)

// ints ints测试
func ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := mysort.IntArray(data) //转换为IntArray
	mysort.Sort(a)
	if !mysort.IsSorted(a) {
		panic("fails")
	}
	fmt.Printf("The sorted array is %v\n", a)
}

// strings stings测试
func strings() {
	data := []string{"monday", "friday", "tuesday", "wednesday", "sunday", "thursday", "", "saturday"}
	a := mysort.StringArray(data)
	mysort.Sort(a)
	if !mysort.IsSorted(a) {
		panic("fails")
	}
	fmt.Printf("The sorted array is %v\n", a)
}

// days days测试
func days() {
	Sunday := mysort.Day{0, "SUN", "Sunday"}
	Monday := mysort.Day{1, "MON", "Monday"}
	Tuesday := mysort.Day{2, "TUE", "Tuesday"}
	Wednesday := mysort.Day{3, "WED", "Wednesday"}
	Thursday := mysort.Day{4, "THU", "Thursday"}
	Friday := mysort.Day{5, "FRI", "Friday"}
	Saturday := mysort.Day{6, "SAT", "Saturday"}
	// data := []*mysort.Day{&Sunday, &Monday, &Tuesday, &Wednesday, &Thursday, &Friday, &Saturday}
	// // 定义并初始化
	// a := mysort.DayArray{data}
	a := mysort.DayArray{[]*mysort.Day{&Sunday, &Monday, &Tuesday, &Wednesday, &Thursday, &Friday, &Saturday}}
	mysort.Sort(&a)
	if !mysort.IsSorted(&a) {
		panic("fails")
	}
	// for _, d := range data {
	// 	fmt.Printf("%s ", d.LongName)
	// }
	fmt.Printf("The sorted array is %v\n", a.Data)
	for _, d := range a.Data {
		fmt.Printf("%v ", *d)
	}
	fmt.Println()
}

func main() {
	// 1.指针方法可以通过指针调用
	// 2.值方法可以通过值调用
	// 3.接收者是值的方法可以通过指针调用，因为指针会首先被解引用
	// 4.接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址

	// 接口是一种契约，实现类型必须满足它，它描述了类型的行为，规定类型可以做什么。
	// 接口彻底将类型能做什么，以及如何做分离开来，使得相同接口的变量在不同的时刻表现
	// 出不同的行为，这就是多态的本质。
	// 编写参数是接口变量的函数，这使得它们更具有一般性。
	// 使用接口使代码更具有普适性。
	ints()
	strings()
	days()
}

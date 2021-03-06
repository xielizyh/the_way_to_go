package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

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
	// 11.1
	fmt.Println("------Example 11.10-----")
	root := NewNode(nil, nil)
	root.SetData("root node")
	left := NewNode(nil, nil)
	left.SetData("left node")
	right := NewNode(nil, nil)
	right.SetData("right node")
	root.left = left
	root.right = right
	fmt.Printf("root: %v\n", root)
	fmt.Printf("left: %p\n", left)
	fmt.Printf("right: %p\n", right)

	// 11.15
	fmt.Println("------Example 11.15-----")
	print(Day(1), "was", Celsius(25.0))

	// 11.12
	// 任何提供了接口方法实现代码的类型都隐式地实现了该接口，而不用显式地声明。
	fmt.Println("------Example 11.15-----")
	task := NewTask("start", log.New(os.Stdout, "[xli]", log.Lshortfile|log.Ldate|log.Ltime))
	task.Printf("hello")

	// 11.18
	fmt.Println("------Example 11.18-----")
	carsTest()
}

// Node 节点
type Node struct {
	left  *Node
	data  interface{}
	right *Node
}

// NewNode 创建节点
func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

// SetData 设置节点数据
func (n *Node) SetData(data interface{}) {
	n.data = data
}

// Stringer 字符串接口
type Stringer interface {
	String() string
}

// Celsius 温度类型
type Celsius float64

// Sting 方法
func (c Celsius) String() string {
	return strconv.FormatFloat(float64(c), 'f', 1, 64) + " °C"
}

// Day 天
type Day int

var dayName = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

// String 方法
func (d Day) String() string {
	return dayName[d]
}

func print(args ...interface{}) {
	for i, arg := range args {
		if i > 0 {
			os.Stdout.WriteString(" ")
		}
		switch a := arg.(type) {
		case Stringer:
			// fmt.Println("here 1")
			os.Stdout.WriteString(a.String())
		case int:
			// fmt.Println("here 2")
			os.Stdout.WriteString(strconv.Itoa(a))
		case string:
			// fmt.Println("here 3")
			os.Stdout.WriteString(a)
		default:
			os.Stdout.WriteString("???")
		}
	}
}

// Task Task包含了Logger类型的指针，那么可以调用该接口的方法
// 当一个类型包含（内嵌）另一个类型（实现了一个或多个接口）的指针时，
// 这个类型就可以使用（另一个类型）所有的接口方法。
type Task struct {
	cmd string
	*log.Logger
}

// NewTask 创建
func NewTask(cmd string, logger *log.Logger) *Task {
	return &Task{cmd, logger}
}

// Any 空接口
type Any interface{}

// Car 车
type Car struct {
	Model        string
	Manufacturer string
	BuildYear    int
}

// Cars 切片
type Cars []*Car

func carsTest() {
	ford := &Car{"Fiesta", "Ford", 2008}
	bmw := &Car{"XL 450", "BMW", 2011}
	merc := &Car{"D600", "Mercedes", 2009}
	bmw2 := &Car{"X 800", "BMW", 2008}

	allCars := Cars([]*Car{ford, bmw, merc, bmw2})
	allNewBMWs := allCars.FindAll(func(car *Car) bool {
		return (car.Manufacturer == "BMW") && (car.BuildYear > 2010)
	})
	fmt.Println("AllCars: ", allCars)
	fmt.Println("All BMWs: ", allNewBMWs)

	manufacturers := []string{"Ford", "Aston Martin", "Land Rover", "BMW", "Jaguar"}
	sortedAppender, sortedCars := MakeSortedAppender(manufacturers)
	fmt.Println("Before map sortedCars:", sortedCars)
	// 闭包（匿名）函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量。
	allCars.Process(sortedAppender)
	fmt.Println("Map sortedCars: ", sortedCars)
	BMWCount := len(sortedCars["BMW"])
	fmt.Println("We have ", BMWCount, " BMWs")
}

// Process 处理
func (cs Cars) Process(f func(car *Car)) {
	for _, c := range cs {
		f(c)
	}
}

// FindAll 寻找匹配车辆
func (cs Cars) FindAll(f func(car *Car) bool) Cars {
	cars := make([]*Car, 0)
	cs.Process(func(c *Car) {
		if f(c) {
			cars = append(cars, c)
		}
	})
	return cars
}

// Map 创建新数据
func (cs Cars) Map(f func(car *Car) Any) []Any {
	result := make([]Any, 0, len(cs))
	ix := 0
	cs.Process(func(c *Car) {
		result[ix] = f(c)
		ix++
	})
	return result
}

// MakeSortedAppender 分类
func MakeSortedAppender(manufacturers []string) (func(c *Car), map[string]Cars) {
	sortedCars := make(map[string]Cars)
	for _, m := range manufacturers {
		sortedCars[m] = make([]*Car, 0)
	}
	sortedCars["Default"] = make([]*Car, 0)
	appender := func(c *Car) {
		if _, ok := sortedCars[c.Manufacturer]; ok {
			sortedCars[c.Manufacturer] = append(sortedCars[c.Manufacturer], c)
		} else {
			sortedCars["Default"] = append(sortedCars["Default"], c)
		}
	}
	return appender, sortedCars
}

package main

import (
	floatsort "excersise/float_sort"
	interfacesext "excersise/interfaces_ext"
	interfacespoly2 "excersise/interfaces_poly2"
	mininterface "excersise/min_interface"
	pointinterfaces "excersise/point_interfaces"
	simpleinterface "excersise/simple_interface"
	simpleinterface3 "excersise/simple_interface3"
	sortpersons "excersise/sort_persons"
	stackgeneral "excersise/stack_general"
	"fmt"
	"math"
)

func main() {
	// 11.1
	fmt.Println("------Excersise 11.1-----")
	// 接口变量是指向某个实例变量的引用
	// 方式一
	// var simple simpleinterface.Simple
	// simpler := simpleinterface.Simpler(&simple) // 注意要用地址"&"，因为接受者类型是指针
	// simpleinterface.SimpleFunc(simpler)
	// 方式二
	// simple := new(simpleinterface.Simple)
	// simple.Set(1)
	// simpler := simpleinterface.Simpler(simple)
	// simpleinterface.SimpleFunc(simpler)
	// 方式三
	simpler := simpleinterface.Simpler(&simpleinterface.Simple{Num: 1})
	simpleinterface.SimpleFunc(simpler)
	if t, ok := simpler.(*simpleinterface.Simple); ok {
		fmt.Printf("The type of areIntf is: %T\n", t)
	}

	// 11.2
	fmt.Println("------Excersise 11.2-----")
	sh := interfacespoly2.Shape{}
	s := new(interfacespoly2.Square)
	s.Set(5)
	s.Shape = sh
	r := new(interfacespoly2.Rectangle)
	r.Set(5, 3)
	r.Shape = sh
	c := new(interfacespoly2.Circle)
	c.Set(5)
	c.Shape = sh
	shapes := []interfacespoly2.Shaper{s, r, c}
	for n := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}

	// 11.4
	fmt.Println("------Excersise 11.4-----")
	fi := func(item simpleinterface.Simpler) {
		switch t := item.(type) {
		case *simpleinterface.Simple:
			fmt.Printf("Type Simple %T with value %v\n", t, t)
		case *simpleinterface.RSimple:
			fmt.Printf("Type RSimple %T with value %v\n", t, t)
		case nil:
			fmt.Printf("Type nil\n")
		default:
			fmt.Printf("Unexpected type %T\n", t)
		}
	}
	fi(simpler)
	var rs simpleinterface.RSimple
	fi(&rs)
	var value INT = 3
	fi(&value)

	// 11.5
	fmt.Println("------Excersise 11.5-----")
	tr := interfacesext.Triangle{3, 5}
	var areaIntf interfacesext.AreaInterface
	areaIntf = &tr
	fmt.Printf("The triangle has area: %f\n", areaIntf.Area())
	periIntf := interfacesext.PeriInterface(&interfacesext.Square{Side: 5})
	fmt.Printf("The square has perimeter: %f\n", periIntf.Perimeter())

	// 11.6
	fmt.Println("------Excersise 11.6-----")
	var m pointinterfaces.Magnitude
	p1 := new(pointinterfaces.Point)
	p1.X = 10
	p1.Y = 10
	m = p1
	fmt.Printf("The length of the vector p1 is: %f\n", m.Abs())
	p2 := &pointinterfaces.Point3{X: 10, Y: 10, Z: 10}
	m = p2
	fmt.Printf("The length of the vector p2 is: %f\n", m.Abs())
	// 注意这里没有加"&"，因为Polar接收者方法是值接收，可加可不加，如果加了，就是自动解引用
	m = pointinterfaces.Polar{2.0, math.Pi / 2.0}
	fmt.Printf("The length of the vector p3 is: %f\n", m.Abs())

	// 11.7
	fmt.Println("------Excersise 11.7-----")
	floatArray := floatsort.NewFloat64Array(25)
	floatArray.Fill(10)
	fmt.Printf("Before sorting: %s\n", floatArray)
	floatsort.Sort(floatArray)
	if !floatsort.IsSorted(floatArray) {
		panic("fails")
	}
	fmt.Printf("After sorting: %s\n", floatArray)

	// 11.8
	fmt.Println("------Excersise 11.8-----")
	// 利用sort包中的接口
	sortPersons := make(sortpersons.Persons, 3)
	sortPersons[0].FirstName = "Xie"
	sortPersons[0].LastName = "Li"
	sortPersons[1].FirstName = "John"
	sortPersons[1].LastName = "Smith"
	sortPersons[2].FirstName = "Xie"
	sortPersons[2].LastName = "Ling"
	sortPersons.Sort()
	if !sortPersons.IsSorted() {
		panic("fails")
	}
	fmt.Printf("The sorted array is: %v\n", sortPersons)

	// 11.9
	fmt.Println("------Excersise 11.9-----")
	var s3 simpleinterface3.Simple = simpleinterface3.Simple{Num: 6}
	fmt.Println(simpleinterface3.GI(&s3))

	// 11.0
	fmt.Println("------Excersise 11.11-----")
	minArray := mininterface.IntArray{2, 5, 7, 1, 4, 0}
	min := mininterface.Min(minArray)
	fmt.Printf("The minimum of the array is: %v\n", min)
	minString := mininterface.StringArray{"ddd", "eee", "bbb", "ccc", "aaa"}
	min = mininterface.Min(minString)
	fmt.Printf("The minimum of the array is: %v\n", min)
	// var temp string = min
	// fmt.Println(minString)
	var num int = 1
	f := func(any interface{}) {
		println(any)
	}
	f(num)

	// 11.11
	fmt.Println("------Excersise 11.11-----")
	// <目标类型的值>，<布尔参数> := <表达式>.( 目标类型 ) // 安全类型断言
	mf := func(i obj) obj {
		switch i.(type) {
		case int:
			return i.(int) * 2
		case string:
			return i.(string) + i.(string)
		case float64:
			return i.(float64) * 2
		}
		return i
	}
	isl := []obj{0, 1, 2, 3, 4, 5}
	res1 := mapFunc(mf, isl)
	for _, v := range res1 {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	ssl := []obj{"a", "b", "c", "d", "e", "f"}
	res2 := mapFunc(mf, ssl)
	for _, v := range res2 {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	fmt.Println("------Excersise 11.12-----")
	res3 := mapFuncVar(mf, isl...)
	for _, v := range res3 {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	res4 := mapFuncVar(mf, ssl...)
	for _, v := range res4 {
		fmt.Printf("%v ", v)
	}
	fmt.Println()
	// 混合型
	msl := []obj{"a", 5, "hello", 2.3}
	res5 := mapFuncVar(mf, msl...)
	for _, v := range res5 {
		fmt.Printf("%v ", v)
	}
	// 11.13
	fmt.Println()
	fmt.Println("------Excersise 11.13-----")
	var st stackgeneral.Stack
	// 为什么没有用&st.Push(1)，因为st本质是切片
	// (&st).Push(1)
	st.Push(1)
	st.Push("hello world")
	st.Push(1.1)
	fmt.Println(st)
	// fmt.Println(st[:3])
	poped, err := st.Pop()
	if err == nil {
		fmt.Println("Pop data is ", poped)
	}

	// 容量够的情况下测试
	stTmp := make(stackgeneral.Stack, 0, 3)
	fmt.Printf("the stTmp is %p\n", stTmp)
	stTmp.PushTmp(1)
	// 注意stTmp切片的长度和容量没有改变，因为本质len,cap两个参数是值传递的
	fmt.Printf("the stTmp is %p\n", stTmp)
	fmt.Printf("The len, cap of stTmp is %d, %d\n", len(stTmp), cap(stTmp))
	fmt.Println(stTmp[:3])
	stTmp.PushTmp("hello world")
	fmt.Println(stTmp[:3])
	stTmp.PushTmp(1.1)

	// var stTmp1 stackgeneral.Stack
	// stTmp1.PushTmp("test")
	// fmt.Printf("the stTmp1 is %p\n", stTmp1)
	// fmt.Println(stTmp)
	// 由于容量够用，因此还是在原始数组上添加的
	fmt.Println(stTmp[:3])
	poped, err = stTmp.Pop()
	if err == nil {
		fmt.Println("Pop data is ", poped)
	}

	fmt.Println("test slice ptr...")
	// 切片的内部结构是一个结构体，包括指针，长度，容量三个字段。其中
	// 指针指向底层数组中切片指定的开始位置。因此切片作为函数参数时其实是传递结构体，也就是值传递，
	// 在函数中对slice的修改是通过slice中保存的地址对底层数组进行修改。
	// 当需要对slice做插入和删除时，由于需要更改长度字段，值拷贝就不行了，需要传slice本身在内存中的地址。
	// 测试切片传参
	sliFunc := func(sli []int, n int) {
		sli[0] = 11
		sli[1] = 22
		// 注意这里改变了切片的长度了，需要传入切片的地址
		sli = append(sli, n)
	}
	slice1 := make([]int, 2)
	sliFunc(slice1, 255)
	fmt.Println(slice1)
	// 测试切片指针传参
	sliPtrFunc := func(pSli *[]int, n int) {
		sli := *pSli
		sli[0] = 11
		sli[1] = 22
		*pSli = append(sli, n)
	}
	slice2 := make([]int, 2)
	sliPtrFunc(&slice2, 255)
	fmt.Println(slice2)
}

// INT INT
type INT int

// Set Set方法
func (i *INT) Set(n int) {
	*i = INT(n)
}

// Get Get方法
func (i *INT) Get() int {
	return int(*i)
}

type obj interface{}

func mapFunc(mf func(obj) obj, list []obj) []obj {
	result := make([]obj, len(list))

	for ix, v := range list {
		result[ix] = mf(v)
	}
	return result
}

func mapFuncVar(mf func(obj) obj, list ...obj) []obj {
	result := make([]obj, len(list))

	for ix, v := range list {
		result[ix] = mf(v)
	}

	return result
}

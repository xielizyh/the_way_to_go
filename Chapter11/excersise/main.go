package main

import (
	floatsort "excersise/float_sort"
	interfacesext "excersise/interfaces_ext"
	interfacespoly2 "excersise/interfaces_poly2"
	pointinterfaces "excersise/point_interfaces"
	simpleinterface "excersise/simple_interface"
	sortpersons "excersise/sort_persons"
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

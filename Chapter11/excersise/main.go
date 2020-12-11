package main

import (
	interfacespoly2 "excersise/interfaces_poly2"
	simpleinterface "excersise/simple_interface"
	"fmt"
)

func main() {
	// 11.1
	fmt.Println("------Excersise 11.1-----")
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

	// 11.1
	fmt.Println("------Excersise 11.1-----")
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

}

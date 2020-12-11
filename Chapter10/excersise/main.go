package main

import (
	"fmt"
	"time"
)

func main() {
	// 10.1
	fmt.Println("------Excersise 10.1------")
	ShowVCard()
	// 10.2
	fmt.Println("------Excersise 10.2------")
	var pers1 Person
	pers1.firstName = "John"
	pers1.lastName = "Smith"
	pers1 = UpPerson(pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)
	// 10.3
	fmt.Println("------Excersise 10.3------")
	p1 := new(Point)
	p1.X = 3
	p1.Y = 4
	fmt.Printf("The length of the vector p1 is: %f\n", Abs(p1))
	p2 := &Point{X: 4, Y: 5}
	fmt.Printf("The length of the vector p2 is: %f\n", Abs(p2))
	q := Scale(p1, 5)
	fmt.Printf("The length of the vector q is: %f\n", Abs(&q))
	fmt.Printf("Point p1 scaled by 5 has the following coordinates: X %f - Y %f\n", q.X, q.Y)
	// 10.4
	fmt.Println("------Excersise 10.4------")
	r := new(Rectangle)
	r.Length = 10
	r.Width = 20
	fmt.Printf("The area and perimeter of the r is: %d %d\n", Area(r), Perimeter(r))
	// 10.5
	fmt.Println("------Excersise 10.5------")
	anony := NewAnonymous(1., 2, "Hello")
	fmt.Printf("anony.n=%f, anony.int=%d, anony.string=%s\n", anony.n, anony.int, anony.string)
	// 10.6
	fmt.Println("------Excersise 10.6------")
	// 类型和作用在它上面定义的方法必须在同一个包里定义,这就是为什么不能在 int、float 或类似这些的类型上定义方法
	// 但是有一个间接的方式：可以先定义该类型（比如：int 或 float）的别名类型，然后再为别名类型定义方法。
	// 或者将它作为匿名类型嵌入在一个新的结构体中。当然方法只在这个别名类型上有效。
	employee := Employee{10000}
	employee.GiveRaise(10)
	fmt.Printf("After raise 10%%, the salary is %d\n", employee.salary)
	// 类型和作用在它上面定义的方法必须在同一个包里定义,这就是为什么不能在 int、float 或类似这些的类型上定义方法
	// 但是有一个间接的方式：可以先定义该类型（比如：int 或 float）的别名类型，然后再为别名类型定义方法。
	// 或者将它作为匿名类型嵌入在一个新的结构体中。当然方法只在这个别名类型上有效。
	m := myTime{time.Now()}
	fmt.Println("Full time now:", m.String())
	fmt.Println("First 3 chars:", m.first3Chars())
	// 函数和方法的区别：函数将变量作为参数，方法在变量上被调用
	// 方法没有和数据定义（结构体）混在一起：它们是正交的类型；表示（数据）和行为（方法）是独立的。
	// 10.8
	fmt.Println("------Excersise 10.8------")
	mc := Mercedes{Car{4, nil}}
	fmt.Printf("The Mercedes has %d wheel\n", mc.numberOfWheels())
	mc.goToWorkIn()
	mc.sayHiToMerker()
	// 10.9
	fmt.Println("------Excersise 10.9------")
	p3 := &PointNew{3, 4}
	fmt.Printf("The length of the vector p3 is %f\n", p3.Abs())
	q3 := p3.Scale(5)
	fmt.Printf("The length of the vector q is: %f\n", q3.Abs())
	fmt.Printf("Point p3 scaled by 5 has the following coordinates: X %f - Y %f\n", q3.X, q3.Y)
	// 10.11
	fmt.Println("------Excersise 10.11------")
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
	// 10.2
	fmt.Println("------Excersise 10.2------")
	var vi integer = integer{10}
	f(vi.n)
	// 10.12
	fmt.Println("------Excersise 10.12------")
	t := &T{7, -2.35, "abc\tdef"}
	fmt.Println(t)
	fmt.Printf("%v\n", t)
	fmt.Println(*t)
	// 10.13
	fmt.Println("------Excersise 10.13------")
	var celsius Celsius = 36.3
	// 注意调用方法的类型要与类型的String方法接收者类型一致才会替换输出
	// 否则仍然是自带库中的string方法
	fmt.Println(celsius)
	fmt.Println(&celsius)
	// 10.14
	fmt.Println("------Excersise 10.14------")
	var day Day = WE
	fmt.Println(day)
	// 10.15
	fmt.Println("------Excersise 10.15------")
	var tz TZ = UTC
	fmt.Println(tz)
	fmt.Println(0 * HOUR)
}

type myTime struct {
	time.Time
}

func (t myTime) first3Chars() string {
	return t.Time.String()[0:3]
}

// Base base
type Base struct{}

// Magic magic
func (Base) Magic() {
	fmt.Println("base magic")
}

// MoreMagic moremagic
func (b Base) MoreMagic() {
	b.Magic()
	b.Magic()
}

// Voodoo voodoo
type Voodoo struct {
	Base
}

// Magic magic
func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

type integer struct {
	n int
}

func (i integer) get() int {
	return i.n * i.n
}

func f(i int) {
	fmt.Println(i)
}

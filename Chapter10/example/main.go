package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// Person 定义人的名字
type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func main() {
	// 10.2
	fmt.Println("------Example 10.2------")
	// 1-struct作为值类型
	var pers1 Person
	pers1.firstName = "John"
	pers1.lastName = "Smith"
	upPerson(&pers1)
	fmt.Printf("The name of the person is %s %s\n", pers1.firstName, pers1.lastName)
	// 2-struct作为指针类型
	pers2 := new(Person)
	pers2.firstName = "John"    //无论结构体类型还是结构体类型指针，都使用同样的选择器符访问结构体字段，即"."
	pers2.lastName = "Smith"    //并没有像C语言中使用"->"访问
	(*pers2).firstName = "John" // 取指针的内容
	upPerson(pers2)
	fmt.Printf("The name of the person is %s %s\n", pers2.firstName, pers2.lastName)
	// 3-struct作为结构体字面量
	pers3 := &Person{"John", "Smith"}
	upPerson(pers3)
	fmt.Printf("The name of the person is %s %s\n", pers3.firstName, pers3.lastName)

	// var identifier [type] = value，type可省略，因为go可以自动推断类型
	var tmp = 3
	fmt.Println(tmp)
	type ps Person
	a := ps{"xie", "li"}
	var b = a
	var c Person = Person(a)
	var d = Person(a)
	fmt.Printf("Type(a)=%T, Type(b)=%T, Type(c)=%T, Type(d)=%T\n", a, b, c, d)
	// 10.7
	fmt.Println("------Example 10.7------")
	tt := tagType{field1: true, field2: "Barak Obama", field3: 1}
	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}
	// 10.8
	fmt.Println("------Example 10.8------")
	// 结构体可以包含一个或多个匿名字段
	// 这些字段没有显式的名字，只有字段的类型是必须的，此时类型就是字段的名字
	type T struct {
		F1 int `json:"f_1"`
		F2 int `json:"f_2,omitempty"`
		// omitempty忽略为空的字段
		F3 int `json:"f_3,omitempty"`
		// 序列化时使用"-"忽略任一字段
		F4 int `json:"-"`
	}
	t := T{1, 0, 2, 3}
	e, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", e) // {"f_1":1,"f_3":2}
	// 10.11
	// Go 方法是作用在接收者（receiver）上的一个函数，接收者是某种类型的变量
	fmt.Println("------Example 10.11------")
	fmt.Println(intVector{1, 2, 3}.sum())
	// 10.14
	fmt.Println("------Example 10.14------")
	// 方法既可以按值调用，也可以在指向这个类型的指针上调用（go会自动解调用）
	// 指针方法和值方法都可以在指针或非指针上被调用
	// 值
	var lst List
	lst.Append(1)
	fmt.Printf("%v (len:%d)\n", lst, lst.Len())
	// 指针
	plst := new(List)
	plst.Append(2)
	fmt.Printf("%v (len:%d)\n", plst, plst.Len())
	// 10.19
	fmt.Println("------Example 10.19------")
	// A：聚合（或组合）：包含一个所需功能类型的具名字段。
	ct1 := new(Customer1)
	ct1.Name = "Barak Obama"
	ct1.log = new(Log1)
	ct1.log.msg = "1 - Yes we can"
	ct1.Log1().Add1("2 - After me the world will be a better place!")
	fmt.Println(ct1.Log1())
	fmt.Println(ct1.log)
	// B：内嵌：内嵌（匿名地）所需功能类型
	ct2 := &Customer2{"Barak Obama", Log2{"1 - Yes we can!"}}
	ct2.Add2("2 - After me the world will be a better place!")
	fmt.Println(ct2)
	fmt.Println(fmt.Sprintf("Hello world!"))
	lg2 := new(Log2)
	lg2.msg = "test string self define"
	fmt.Println(lg2)
	// 不要在 String() 方法里面调用涉及 String() 方法的方法，它导致了一个无限递归调用
	// 10.8
	fmt.Println("------Example 10.8------")
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d Kb\n", m.Alloc/1024)
}

type tagType struct {
	field1 bool   `json: "field1"`
	field2 string `json: "field2"`
	field3 int    `json: "field3"`
}

func refTag(tt tagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}

type intVector []int

func (v intVector) sum() (s int) {
	for _, v := range v {
		s += v
	}
	return
}

// List 类型
type List []int

// Len 类型List值上方法
func (l List) Len() int { return len(l) }

// Append 类型List指针上方法
func (l *List) Append(val int) { *l = append(*l, val) }

// Log1 日志
type Log1 struct {
	msg string
}

// Customer1 个人化
type Customer1 struct {
	Name string
	log  *Log1
}

// Add1 增加日志
func (l *Log1) Add1(s string) {
	l.msg += "\n" + s
}

// String 返回日志
func (l *Log1) String() string {
	return l.msg
}

// Log1 日志
func (c *Customer1) Log1() *Log1 {
	return c.log
}

// Log2 日志
type Log2 struct {
	msg string
}

// Customer2 个人化
type Customer2 struct {
	Name string
	Log2
}

// Add2 增加日志
func (l *Log2) Add2(s string) {
	l.msg += "\n" + s
}

// String 返回日志
func (l *Log2) String() string {
	// fmt.Println("test", l.msg)
	return l.msg
}

// String 字符串方法
// golang允许你通过实现 String() 接口来自定义输出
// 那么它会被用在fmt.Printf() 中生成默认的输出：
// 等同于使用格式化描述符 %v 产生的输出。
// 还有 fmt.Print() 和 fmt.Println() 也会自动使用 String() 方法
func (c *Customer2) String() string {
	// c.Log2为什么没有调用Log2的String接口呢？
	return c.Name + "\nLog:" + fmt.Sprintln(c.Log2)
}

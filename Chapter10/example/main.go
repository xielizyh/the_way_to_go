package main

import (
	"encoding/json"
	"fmt"
	"reflect"
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

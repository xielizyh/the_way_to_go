package main

import "strings"

// Person 定义人的名字
type Person struct {
	firstName string
	lastName  string
}

// UpPerson 名字全大写
func UpPerson(p Person) Person {
	return Person{strings.ToUpper(p.firstName), strings.ToUpper(p.lastName)}
}

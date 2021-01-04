package main

import (
	"fmt"
	"strconv"
)

// Person 人
type Person struct {
	Name   string
	salary float64
	chF    chan func()
}

// NewPerson 创建
func NewPerson(name string, salary float64) *Person {
	p := &Person{name, salary, make(chan func())}
	go p.backend()
	return p
}

func (p *Person) backend() {
	for f := range p.chF {
		f()
	}
}

// SetSalary 设置工资
func (p *Person) SetSalary(salary float64) {
	p.chF <- func() { p.salary = salary }
}

// Salary 检索工资
func (p *Person) Salary() float64 {
	fChan := make(chan float64)
	p.chF <- func() { fChan <- p.salary }
	return <-fChan
}

// String string
func (p *Person) String() string {
	return "Person - name is: " + p.Name + " - salary is: " + strconv.FormatFloat(p.Salary(), 'f', 2, 64)
}

func concAccess() {
	bs := NewPerson("Smith Bill", 2500.5)
	fmt.Println(bs)
	bs.SetSalary(4000.25)
	fmt.Println("Salary changed: ")
	fmt.Println(bs)
}

package main

// Employee 雇员
type Employee struct {
	salary int
}

// GiveRaise 涨工资
func (e *Employee) GiveRaise(percent int) {
	e.salary += e.salary * percent / 100
}

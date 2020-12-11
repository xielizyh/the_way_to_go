package main

import (
	"fmt"
	"point_methods/employee"
	"point_methods/hello"
)

func main() {
	// 一定注意要加包名啊!!!
	hello.Hello()
	employee := new(employee.Employee)
	employee.Salary = 10000
	employee.FirstName = "John"
	employee.LastName = "Smith"
	employee.SetID(521)
	fmt.Println(employee.ID())
}

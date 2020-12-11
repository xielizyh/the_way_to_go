package employee

// Base 基本
type Base struct {
	id int
}

// Person 人员
type Person struct {
	FirstName string
	LastName  string
	Base
}

// Employee 雇员
type Employee struct {
	Salary int
	Person
}

// ID 返回id
func (b *Base) ID() int { return b.id }

// SetID 修改ID
func (b *Base) SetID(id int) {
	b.id = id
}

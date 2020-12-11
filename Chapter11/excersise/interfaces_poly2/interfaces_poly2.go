package interfacespoly2

import "math"

// Shaper 形状
type Shaper interface {
	Area() float32
}

// Square 正方形
type Square struct {
	side float32
	Shape
}

// Set 设置
func (s *Square) Set(side float32) {
	s.side = side
}

// Area 正方形面积
func (s *Square) Area() float32 {
	return s.side * s.side
}

// Rectangle 长方形
type Rectangle struct {
	length, width float32
	Shape
}

// Set 设置
func (r *Rectangle) Set(length, width float32) {
	r.length = length
	r.width = width
}

// Area 长方形面积
func (r *Rectangle) Area() float32 {
	return r.length * r.width
}

// Circle 圆形
type Circle struct {
	radius float32
	Shape
}

// Set 设置
func (c *Circle) Set(radius float32) {
	c.radius = radius
}

// Area 圆形面积
func (c *Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

// Shape 抽象类型
type Shape struct{}

// Area 抽象类型面积
func (s *Shape) Area() float32 {
	return -1
}

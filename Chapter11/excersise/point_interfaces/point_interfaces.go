package pointinterfaces

import "math"

// Magnitude 接口
type Magnitude interface {
	Abs() float64
}

// Point 二维坐标
type Point struct {
	X, Y float64
}

// Point3 三维坐标
type Point3 struct {
	X, Y, Z float64
}

// Polar 极坐标
type Polar struct {
	R, Theta float64
}

// Abs 方法
func (p *Point) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

// Abs 方法
func (p *Point3) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y + p.Z*p.Z))
}

// Abs 方法
func (p Polar) Abs() float64 {
	return p.R
}

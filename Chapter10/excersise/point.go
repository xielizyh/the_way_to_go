package main

import "math"

// Point 点坐标
type Point struct {
	X float64
	Y float64
}

// Point3 三维坐标
type Point3 struct {
	X, Y, Z float64
}

// Polar 极坐标
type Polar struct {
	R, T float64
}

// Abs 求坐标向量长度
func Abs(p *Point) float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

// Scale 尺度
func Scale(p *Point, s float64) (q Point) {
	q.X = p.X * s
	q.Y = p.Y * s
	return
}

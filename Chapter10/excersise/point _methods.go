package main

import "math"

// PointNew 点坐标
type PointNew struct {
	X float64
	Y float64
}

// Point3New 三维坐标
type Point3New struct {
	X, Y, Z float64
}

// PolarNew 极坐标
type PolarNew struct {
	R, T float64
}

// Abs 求坐标向量长度
func (p *PointNew) Abs() float64 {
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

// Scale 尺度
func (p *PointNew) Scale(s float64) (q PointNew) {
	q.X = p.X * s
	q.Y = p.Y * s
	return
}

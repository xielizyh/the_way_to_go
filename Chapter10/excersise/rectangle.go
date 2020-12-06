package main

// Rectangle 长方形
type Rectangle struct {
	Length int
	Width  int
}

// Area 面积
func Area(r *Rectangle) int {
	return r.Length * r.Width
}

// Perimeter 周长
func Perimeter(r *Rectangle) int {
	return 2 * (r.Length + r.Width)
}

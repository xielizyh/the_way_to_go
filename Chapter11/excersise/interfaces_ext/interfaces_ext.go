package interfacesext

// AreaInterface 面积接口
type AreaInterface interface {
	Area() float32
}

// PeriInterface 周长接口
type PeriInterface interface {
	Perimeter() float32
}

// Triangle 三角形
type Triangle struct {
	Base   float32
	Height float32
}

// Area 面积
func (t *Triangle) Area() float32 {
	return t.Base * t.Height / 2
}

// Square 正方形
type Square struct {
	Side float32
}

// Area 面积
func (s *Square) Area() float32 {
	return s.Side * s.Side
}

// Perimeter 周长
func (s *Square) Perimeter() float32 {
	return 4 * s.Side
}

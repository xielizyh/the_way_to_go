package main

type anonymous struct {
	n float32
	int
	string
}

// NewAnonymous 结构体工厂
func NewAnonymous(f float32, i int, s string) *anonymous {
	anony := &anonymous{f, i, s}

	return anony
}

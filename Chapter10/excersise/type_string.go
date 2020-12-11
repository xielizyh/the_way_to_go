package main

import "fmt"

// T 测试类型的string方法
type T struct {
	a int
	b float32
	c string
}

// String 类型的T的string
func (t *T) String() string {
	// return strconv.Itoa(t.a) + " / " +
	// 	strconv.FormatFloat(float64(t.b), 'f', 6, 64) + " / " + t.c
	return fmt.Sprintf("%d / %f / %q", t.a, t.b, t.c)
	// return fmt.Sprintf("%d / %f / %s", t.a, t.b, t.c)
	// return fmt.Sprintf("%v", t) //会引起递归调用崩溃
}

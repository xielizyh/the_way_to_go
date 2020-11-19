package main

// Adder 将函数作为返回值
func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}

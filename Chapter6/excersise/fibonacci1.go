package main

// Fibonacci1 斐波那契数列用闭包实现
func Fibonacci1() func(int) int {
	var value1 int
	var value2 int

	return func(i int) int {
		if i == 0 {
			value1, value2 = 0, 0
		} else if i == 1 {
			value1, value2 = 0, 1
		}
		value := value1 + value2
		value1, value2 = value2, value
		return value
	}
}

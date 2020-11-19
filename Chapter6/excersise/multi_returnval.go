package main

// Calculate1 计算两个整数的和差积
func Calculate1(a, b int) (int, int, int) {
	add := a + b
	subtract := a - b
	multiply := a * b

	return add, subtract, multiply
}

// Calculate2 计算两个整数的和差积
func Calculate2(a, b int) (add int, subtract int, multiply int) {
	add = a + b
	subtract = a - b
	multiply = a * b

	return
}

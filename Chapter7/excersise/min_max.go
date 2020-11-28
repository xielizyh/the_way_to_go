package main

// MinSlice 返回切片的最小值
func MinSlice(num ...int) int {
	min := 0
	for _, n := range num {
		if n < min {
			min = n
		}
	}
	return min
}

// MaxSlice 返回切片的最大值
func MaxSlice(num ...int) int {
	max := 0
	for _, n := range num {
		if n > max {
			max = n
		}
	}
	return max
}

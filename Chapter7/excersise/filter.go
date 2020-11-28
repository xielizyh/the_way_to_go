package main

// IsEven 判断一个数是否偶数
func IsEven(x int) bool {
	if x%2 == 0 {
		return true
	}
	return false
}

// Filter 过滤
func Filter(s []int, fn func(int) bool) []int {
	slice := make([]int, 0)
	for _, v := range s {
		if fn(v) {
			slice = append(slice, v)
		}
	}
	return slice[:]
}

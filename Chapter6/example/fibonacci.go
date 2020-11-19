package main

// LIM 最大
const LIM = 41

var fibs [LIM]uint64

// Fibonacci 斐波那契数列
func Fibonacci(n int) (res uint64) {
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}

	if n <= 1 {
		res = 1
	} else {
		res = Fibonacci(n-1) + Fibonacci(n-2)
	}
	fibs[n] = res

	return
}

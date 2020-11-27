package main

// FibonacciFuncArray Fibonacci数组
func FibonacciFuncArray(n int) []int {
	var array [50]int
	// array := make([]int, n)
	FibonacciArray(&array)

	return array[:n]
}

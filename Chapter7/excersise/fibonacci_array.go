package main

// FibonacciArray Fibonacci数组
func FibonacciArray(arr *[50]int) {
	for i := range arr {
		if i <= 1 {
			arr[i] = 1
		} else {
			arr[i] = arr[i-1] + arr[i-2]
		}
	}
}

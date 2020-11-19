package main

import "fmt"

// RecursionPrint 递归打印
func RecursionPrint(n int) {
	if n <= 1 {
		fmt.Printf("%d", n)
	} else {
		fmt.Printf("%d ", n)
		RecursionPrint(n - 1)
	}
}

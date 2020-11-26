package main

import "fmt"

func main() {
	// 7.1
	fmt.Println("------Excersise 7.1------")
	var arr1 [3]int
	ArrayValue(arr1)
	for i := range arr1 {
		fmt.Printf("%d ", arr1[i])
	}
	// 7.2
	fmt.Println("------Excersise 7.2------")
	var arr2 [50]int
	FibonacciArray(&arr2)
	for i := range arr2 {
		fmt.Printf("%d ", arr2[i])
	}
}

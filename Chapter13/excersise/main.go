package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("------Excersise 13.1-----")
	// a, b := 10, 0
	// quotient, err := recoverDividebyZero(a, b)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%d / %d = %.2f\n", a, b, quotient)
	fmt.Println("------Excersise 13.3-----")
	var num1 int64 = math.MaxInt32 + 1
	num2, err := IntFromInt64(num1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Convert int64 %d to int %d\n", num1, num2)
}

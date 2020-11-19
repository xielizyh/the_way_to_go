package main

import "fmt"

func main() {
	var number int
	fmt.Println("Please input a number:")
	fmt.Scanln(&number)
	switch number {
	case 1, 2, 3:
		fmt.Println("It is Q1th")
	case 4, 5, 6:
		fmt.Println("It is Q2nd")
	case 7, 8, 9:
		fmt.Println("It is Q3rd")
	case 10, 11, 12:
		fmt.Println("It is Q4th")
	default:
		fmt.Println("Please check your input")
	}

}

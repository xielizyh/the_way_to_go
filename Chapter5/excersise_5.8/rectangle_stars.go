package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			switch {
			case j == 0, j == 19:
				fmt.Printf("*")
			case i == 0, i == 9:
				fmt.Printf("*")
			default:
				fmt.Printf(" ")
			}
		}
		fmt.Println("")
	}
}

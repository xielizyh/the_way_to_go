package main

import "fmt"

func main() {
	/* 	row := 0
	   	for {
	   		for i := 0; i < row; i++ {
	   			fmt.Printf("G")
	   		}
	   		fmt.Println()
	   		row++
	   		if row > 25 {
	   			break
	   		}
	   	} */

	for i := 0; i < 25; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("G")
		}
		fmt.Println()
	}

	str := "G"
	for i := 0; i < 25; i++ {
		fmt.Println(str)
		str += "G"
	}
}

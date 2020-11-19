package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "124"
	var newS string

	fmt.Printf("The size of ints is %d\n", strconv.IntSize)
	an, err := strconv.Atoi(orig)
	if err != nil {
		fmt.Printf("oring %s is not an integer - exiting with error\n", orig)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	an += 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is %s\n", newS)
}

package main

import (
	"fmt"

	"4.6/trans"
)

var twoPi = 2 * trans.Pi

var a string

func main() {
	fmt.Printf("2*Pi = %g\n", twoPi)

	a = "G"
	print(a)
	f1()
}

func f1() {
	a := "O"
	print(a)
	f2()
}

func f2() {
	print(a)
}

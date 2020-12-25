package main

import (
	"fmt"
	"string_reverse/reverse"
)

func main() {
	src := "This is my test"
	target := reverse.StringReverse(src)
	fmt.Println(target)
}

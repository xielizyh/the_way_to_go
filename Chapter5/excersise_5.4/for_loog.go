package main

import "fmt"

func main() {
	for i := 0; i < 15; i++ {
		fmt.Printf("i=%d\n", i)
	}

	j := 0
loop:
	fmt.Printf("j=%d\n", j)
	j++
	if j < 15 {
		goto loop
	}
}

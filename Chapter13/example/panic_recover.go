package main

import "fmt"

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Panicing %s\n", err)
		}
	}()
	badCall()
	fmt.Printf("After bad call\r\n")
}

// defer-panic-recover
func panicRecover() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}

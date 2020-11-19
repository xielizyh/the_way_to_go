package main

import (
	"fmt"
)

// VarArgs 可变参数
func VarArgs(args ...string) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

package main

import (
	"fmt"
	"os"
	"strings"
)

// helloWho 你好
func helloWho() {
	who := "Hello "
	if len(os.Args) > 1 {
		who += strings.Join(os.Args[1:], " ")
		who += "!"
		fmt.Println(who)
	}
}

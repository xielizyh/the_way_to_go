// 统计
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your message, type S to stop:")
	nrchars, nrwords, nrlines := 0, 0, 0
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred while reading: %s\n", err)
		}
		if input == "S\r\n" {
			fmt.Println("Here are the counts:")
			fmt.Printf("Number of characters: %d\n", nrchars)
			fmt.Printf("Number of nrwords: %d\n", nrwords)
			fmt.Printf("Number of lines: %d\n", nrlines)
			break
		}
		nrchars += len(input) - 2
		nrwords += len(strings.Fields(input))
		//fmt.Println(strings.Fields(input))
		nrlines++
	}
}

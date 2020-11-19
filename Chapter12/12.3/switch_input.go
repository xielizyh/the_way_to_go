package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name:")
	input, err := inputReader.ReadString('\n')

	if err != nil {
		fmt.Println("There were errors reading, exiting program.")
		return
	}

	fmt.Printf("Your name is %s\n", input)
	// For Unix: test with delimter "\n", for windows: test with "\r\n"
	// version 1
	switch input {
	case "Philip\r\n":
		fmt.Println("Welcome Philip")
	case "Ivo\r\n":
		fmt.Println("Welcome Ivo")
	case "Chris\r\n":
		fmt.Println("Welcome Chris")
	default:
		fmt.Println("You are not welcome here! Goodbye.")
	}
	// version 2
	switch input {
	case "Philip\r\n":
		fallthrough
	case "Ivo\r\n":
		fallthrough
	case "Chris\r\n":
		fmt.Printf("Welcome %s\r\n", input)
	default:
		fmt.Println("You are not welcome here! Goodbye.")
	}
	// version 3
	switch input {
	case "Philip\r\n", "Ivo\r\n":
		fmt.Printf("Welcome %s\r\n", input)
	default:
		fmt.Println("You are not welcome here! Goodbye.")
	}
}

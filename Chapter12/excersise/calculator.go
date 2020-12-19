package main

import (
	"bufio"
	stackgeneral "excersise/stack_general"
	"fmt"
	"os"
	"strconv"
)

func calculator() {
	buf := bufio.NewReader(os.Stdin)
	calc1 := new(stackgeneral.Stack)
	fmt.Println("Give a number, an operator (+, -, *, /), or q to stop:")
	for {
		token, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println("Input error:", err)
			os.Exit(1)
		}
		token = token[:len(token)-2] //去除"\r\n"
		switch {
		case token == "q":
			fmt.Println("Calculator stopped")
			return
		case token >= "0" && token <= "999999":
			i, _ := strconv.Atoi(token)
			calc1.Push(i)
		case token == "+":
			p, _ := calc1.Pop()
			q, _ := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)+q.(int))
		case token == "-":
			p, _ := calc1.Pop()
			q, _ := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)-q.(int))
		case token == "*":
			p, _ := calc1.Pop()
			q, _ := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)*q.(int))
		case token == "/":
			p, _ := calc1.Pop()
			q, _ := calc1.Pop()
			fmt.Printf("The result of %d %s %d = %d\n", p, token, q, p.(int)/q.(int))
		default:
			fmt.Println("Invalid input")
		}
	}
	/* 	inputReader := bufio.NewReader(os.Stdin)
	   	var stack stackgeneral.Stack
	   	fmt.Println("Please input expression, type q to stop:")
	   	for {
	   		input, err := inputReader.ReadString('\n')
	   		if err != nil {
	   			fmt.Printf("An error occurred while reading: %s\n", err)
	   		}
	   		if input == "q\r\n" {
	   			fmt.Println("Exit...")
	   			break
	   		}
	   		input = input[:len(input)-2]
	   		// fmt.Println(input)
	   		stack.Push(input)
	   		if stack.Len() == 3 {
	   			num1, num2, result := 0, 0, 0
	   			operator := ""
	   			var any interface{}
	   			any, err = stack.Pop()
	   			if err == nil {
	   				if v, ok := any.(string); ok {
	   					operator = v
	   				}
	   			}
	   			any, err = stack.Pop()
	   			if err == nil {
	   				if v, ok := any.(string); ok {
	   					num2, _ = strconv.Atoi(v)
	   				}
	   			}
	   			any, err = stack.Pop()
	   			if err == nil {
	   				if v, ok := any.(string); ok {
	   					num1, _ = strconv.Atoi(v)
	   				}
	   			}
	   			switch operator {
	   			case "+":
	   				result = num1 + num2
	   			case "-":
	   				result = num1 - num2
	   			case "*":
	   				result = num1 * num2
	   			case "/":
	   				result = num1 / num2
	   			default:
	   				fmt.Println("[ERROR]Operator is ", operator)
	   			}
	   			fmt.Printf("%d %s %d =  %d\n", num1, operator, num2, result)
	   		}
	   	} */
}

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func remove3Till5Char() {
	inputFile, _ := os.Open("test.txt")
	outputFile, _ := os.OpenFile("testT.txt", os.O_WRONLY|os.O_CREATE, 0666)
	defer inputFile.Close()
	defer outputFile.Close()
	inputReader := bufio.NewReader(inputFile)
	outputWriter := bufio.NewWriter(outputFile)
	var outputString string
	for {
		inputString, _, readerError := inputReader.ReadLine()
		if readerError == io.EOF {
			fmt.Println("EOF")
			break
		}
		if len(inputString) < 3 {
			outputString = "\r\n"
		} else if len(inputString) < 5 {
			outputString = string(inputString[2:len(inputString)]) + "\r\n"
		} else {
			outputString = string(inputString[2:5]) + "\r\n"
		}
		_, err := outputWriter.WriteString(outputString)
		if err != nil {
			fmt.Println(err)
			break
		}
	}
	outputWriter.Flush()
}

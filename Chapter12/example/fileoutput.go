package main

import (
	"bufio"
	"fmt"
	"os"
)

func fileoutput() {
	// 追加写 os.O_APPEND
	outputFile, outputError := os.OpenFile("output.dat", os.O_WRONLY|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred while openning or creating\n")
		return
	}
	defer outputFile.Close()
	outputWriter := bufio.NewWriter(outputFile)
	for i := 0; i < 10; i++ {
		outputWriter.WriteString("hello world!\n")
		// outputWriter.Write([]byte("你好世界!\n"))
	}
	outputWriter.Flush()
}

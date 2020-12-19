package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type product struct {
	title    string
	price    float64
	quantity int64
}

func readCSV() {
	inputFile, inputError := os.Open("products.txt")
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}
	defer inputFile.Close()
	// 一定注意啊，切片要用make初始化
	// prod := new([]product)
	prod := make([]product, 4)
	inputReader := bufio.NewReader(inputFile)
	row := 0
	for {
		inputString, readError := inputReader.ReadString('\n')
		spl := strings.Split(inputString, ";")
		// fmt.Println(spl[0])
		prod[row].title = spl[0]
		prod[row].price, _ = strconv.ParseFloat(spl[1], 64)
		prod[row].quantity, _ = strconv.ParseInt(spl[2], 10, 64)
		row++
		if readError == io.EOF {
			break
		}
	}

	for _, value := range prod {
		// fmt.Printf("title:%s, price:%.2f, quantity:%d\n", value.title, value.price, value.quantity)
		fmt.Println(value)
	}
}

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func httpFetch2() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your url:")
	input, _ := inputReader.ReadString('\n')
	// url := "http://" + input[:len(input)-2]
	url := "http://" + strings.TrimSuffix(input, "\r\n")
	res, err := http.Get(url)
	checkError1(err)
	data, err := ioutil.ReadAll(res.Body)
	checkError1(err)
	fmt.Printf("Got: %q", string(data))
}

func checkError1(err error) {
	if err != nil {
		log.Fatalf("Get : %v", err)
	}
}

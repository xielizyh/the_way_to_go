package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func httpFetch() {
	res, err := http.Get("http://www.baidu.com")
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

package main

import (
	"fmt"
	"net/http"
)

var urls = []string{
	"https://www.baidu.com",
	"http://www.baidu.com",
}

func pollURL() {
	for _, url := range urls {
		resp, err := http.Head(url)
		if err != nil {
			fmt.Println("Error:", url, err)
		}
		fmt.Println(url, ": ", resp.Status)
	}
}

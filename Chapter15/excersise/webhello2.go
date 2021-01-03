package main

import (
	"fmt"
	"net/http"
	"strings"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// 获取/hello/后的内容
	remPartOfURL := r.URL.Path[len("/hello/"):]
	fmt.Fprintf(w, "Hello %s!", remPartOfURL)
}

func shoutHelloHandler(w http.ResponseWriter, r *http.Request) {
	// 获取/shouthello/后的内容
	remPartOfURL := r.URL.Path[len("/shouthello/"):]
	fmt.Fprintf(w, "Hello %s!", strings.ToUpper(remPartOfURL))
}

func webHello2() {
	http.HandleFunc("/hello/", helloHandler)
	http.HandleFunc("/shouthello/", shoutHelloHandler)
	http.ListenAndServe(":9999", nil)
}

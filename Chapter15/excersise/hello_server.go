package main

import (
	"fmt"
	"net/http"
)

type hello struct{}

// ServeHTTP 实现http.handler接口
func (h hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func helloServer() {
	var h hello
	http.ListenAndServe(":4000", h)
}

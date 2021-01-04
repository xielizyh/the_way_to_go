package main

import (
	"io"
	"log"
	"net/http"
)

const form1 = `
	<html><body>
		<form action="#" method="post" name="bar">
			<input type="text" name="in" />
			<input type="submit" value="submit"/>
		</form>
	</body></html>
`

type handleFnc func(w http.ResponseWriter, r *http.Request)

func simpleServer1(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>hello, world</h1>")
}

func formServer1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case "GET":
		io.WriteString(w, form1)
	case "POST":
		io.WriteString(w, r.FormValue("in"))
	}
}

func robustWebserver() {
	http.HandleFunc("/test1", logPanics(simpleServer1))
	http.HandleFunc("/test2", logPanics(formServer1))
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}

// 包装了对处理函数的调用，可以恢复panic调用
func logPanics(function handleFnc) handleFnc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", r.RemoteAddr, x)
			}
		}()
		function(w, r)
	}
}

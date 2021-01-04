package main

import (
	"io"
	"net/http"
)

const form = `
	<html><body>
		<form action="#" method="post" name="bar">
			<input type="text" name="in" />
			<input type="submit" value="submit"/>
		</form>
	</body></html>
`

func simpleServer(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>hello, world</h1>")
}

func formServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch r.Method {
	case "GET":
		io.WriteString(w, form)
	case "POST":
		io.WriteString(w, r.FormValue("in"))
	}
}

func simpleWebserver() {
	http.HandleFunc("/test1", simpleServer)
	http.HandleFunc("/test2", formServer)
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}

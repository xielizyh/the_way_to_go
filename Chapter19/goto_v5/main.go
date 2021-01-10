package main

import (
	"flag"
	"fmt"
	store "goto_v5/store"
	"net/http"
	"net/rpc"
)

var (
	listenAddr = flag.String("http", ":8080", "http listen address")
	// dataFile   = flag.String("file", "store.gob", "data store file name")
	dataFile   = flag.String("file", "store.json", "data store file")
	hostname   = flag.String("host", "localhost:8080", "host name and port")
	rpcEnabled = flag.Bool("rpc", false, "enable RPC server")
)

// AddForm 表单
const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`

var st *store.URLStore

func main() {
	flag.Parse()
	st = store.NewURLStore(*dataFile)
	if *rpcEnabled {
		rpc.RegisterName("Store", st)
		rpc.HandleHTTP()
	}
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(*listenAddr, nil)
}

// Redirect 重定向
func Redirect(w http.ResponseWriter, r *http.Request) {
	// 获取键值
	key := r.URL.Path[1:]
	fmt.Println("key is ", key)
	// 提前长URL
	var url string
	if err := store.Get(&key, &url); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 重定向至长URL
	http.Redirect(w, r, url, http.StatusFound)
}

// Add 增加
func Add(w http.ResponseWriter, r *http.Request) {
	// 读取URL
	url := r.FormValue("url")
	fmt.Println("req url is ", url)
	if url == "" {
		// 发送标头
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, AddForm)
		return
	}
	// 存储长URL
	var key string
	if err := store.Put(&url, &key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServer)
		return
	}
	// 发送短URL给用户
	fmt.Fprintf(w, "http://%s/%s", *hostname, key)
}

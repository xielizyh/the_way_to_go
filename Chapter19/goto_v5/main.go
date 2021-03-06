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
	dataFile   = flag.String("file", "store.gob", "data store file name")
	// dataFile   = flag.String("file", "store.json", "data store file")
	hostname   = flag.String("host", "localhost:8080", "host name and port")
	masterAddr = flag.String("master", "", "RPC master address")
	rpcEnabled = flag.Bool("rpc", false, "enable RPC server")
)

// AddForm 表单
const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`

// var st *store.URLStore
// 接口
var st store.Store

func main() {
	flag.Parse()
	if *masterAddr != "" { //we are a slave
		st = store.NewProxyStore(*masterAddr)
	} else { // we are the master
		st = store.NewURLStore(*dataFile)
	}
	if *rpcEnabled {
		// rpc.RegisterName是指定特定的名字
		rpc.RegisterName("Store1", st)
		// 注意rpc.Register默认名字为对象的类型名字，即Store
		// https://studygolang.com/articles/8738
		// rpc.Register(st)
		rpc.HandleHTTP()
	}
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	// fmt.Println(*listenAddr)
	err := http.ListenAndServe(*listenAddr, nil)
	fmt.Println(err)
	fmt.Println("end and exiting")
}

// Redirect 重定向
func Redirect(w http.ResponseWriter, r *http.Request) {
	// 获取键值
	key := r.URL.Path[1:]
	fmt.Println("key is ", key)
	// 提前长URL
	var url string
	if err := st.Get(&key, &url); err != nil {
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
	if err := st.Put(&url, &key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// 发送短URL给用户
	fmt.Fprintf(w, "http://%s/%s", *hostname, key)
}

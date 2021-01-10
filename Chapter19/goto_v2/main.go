package main

import (
	"fmt"
	store "goto_v2/store"
	"net/http"
)

// AddForm 表单
const AddForm = `
<form method="POST" action="/add">
URL: <input type="text" name="url">
<input type="submit" value="Add">
</form>
`

var st = store.NewURLStore("store.gob")

func main() {
	http.HandleFunc("/", Redirect)
	http.HandleFunc("/add", Add)
	http.ListenAndServe(":8080", nil)
}

// Redirect 重定向
func Redirect(w http.ResponseWriter, r *http.Request) {
	// 获取键值
	key := r.URL.Path[1:]
	fmt.Println("key is ", key)
	// 提前长URL
	url := st.Get(key)
	fmt.Println("long url is ", url)
	if url == "" {
		http.NotFound(w, r)
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
	key := st.Put(url)
	// 发送短URL给用户
	fmt.Fprintf(w, "http://localhost:8080/%s", key)
}

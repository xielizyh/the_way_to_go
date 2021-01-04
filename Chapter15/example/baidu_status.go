package main

import (
	"encoding/xml"
)

// Status 状态
type Status struct {
	Text string
}

// User 用户
type User struct {
	XMLName xml.Name
	Status  Status
}

func baiduStatus() {
	// resp, _ := http.Get("https://www.baidu.com")
	// user := User{xml.Name{"", "user"}, Status{""}}
	// xml.Unmarshal(resp.Body, &user)
	// fmt.Printf("status: %s\n", user.Status.Text)
}

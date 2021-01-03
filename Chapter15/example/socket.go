package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
)

func socket() {
	var (
		host          = "www.baidu.com"
		port          = 80
		remote        = host + ":" + strconv.Itoa(port)
		msg    string = "GET / \n"
		data          = make([]uint8, 4096)
		read          = true
		count         = 0
	)

	// 创建一个socket
	conn, err := net.Dial("tcp", remote)
	// 发送一个HTTP GET请求
	io.WriteString(conn, msg)
	for read {
		count, err = conn.Read(data)
		read = (err == nil)
		fmt.Printf(string(data[:count]))
	}
	conn.Close()

	/*
		resp, err := http.Get("http://www.baidu.com")
		if err != nil {
			fmt.Println("http.Get err = ", err)
			return
		}

		//爬取的内容在Body，结束时应该关闭
		defer resp.Body.Close()
		fmt.Println("Status = ", resp.Status)
		fmt.Println("StatusCode = ", resp.StatusCode)
		fmt.Println("Header = ", resp.Header)

		// Body是一个IO，需要读取
		//1.定义一个切片存储
		buf := make([]byte, 1024*4)
		//定义一个字符串
		var tmp string
		for {
			//将Body存储进 n
			n, err := resp.Body.Read(buf)
			//如果n=0说明读取完毕
			if n == 0 {
				//打印结束条件
				fmt.Println("read err = ", err)
				break
			}
			//进行字符串累加
			tmp += string(buf[:n])
		}
		//打印字符串（Body）
		fmt.Println("tmp = ", tmp)
	*/
}

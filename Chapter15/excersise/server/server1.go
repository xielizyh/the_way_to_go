package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

var mapUsers map[string]int

func main() {
	mapUsers = make(map[string]int)
	fmt.Println("Starting the server...")
	// 创建listener
	listener, err := net.Listen("tcp", "localhost:50000")
	checkError(err)
	// 监听并接受来自客户端的连接
	for {
		conn, err := listener.Accept()
		checkError(err)
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn) {
	clName := ""
	for {
		buf := make([]byte, 1024)
		len, err := conn.Read(buf)
		// checkError(err)
		if err != nil {
			mapUsers[clName] = 0
			break
		}
		input := string(buf[:len])
		if strings.Contains(input, ": SH") {
			fmt.Println("Server shutting down")
			os.Exit(0)
		}

		if strings.Contains(input, ": WHO") {
			displayList()
		}

		ix := strings.Index(input, "says:")
		clName = input[:ix-1]
		mapUsers[string(clName)] = 1
		fmt.Printf("Received data: %v\n", string(buf[:len]))
	}
}

func checkError(err error) {
	if err != nil {
		panic("Error: " + err.Error())
	}
}

func displayList() {
	fmt.Println("--------------------------------------------")
	fmt.Println("This is the client list: 1=active, 0=inactive")
	for key, value := range mapUsers {
		fmt.Printf("User %s is %d\n", key, value)
	}
	fmt.Println("--------------------------------------------")
}

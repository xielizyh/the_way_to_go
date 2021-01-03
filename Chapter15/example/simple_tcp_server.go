package main

import (
	"flag"
	"fmt"
	"net"
	"strconv"
	"syscall"
)

// 读取的最大长度
const maxRead = 25

func simpleTCPServer() {
	// 解析参数
	flag.Parse()
	// 检测参数
	if flag.NArg() != 2 {
		panic("usage: host port")
	}
	// 参数格式化字符串
	hostAndPort := fmt.Sprintf("%s:%s", flag.Arg(0), flag.Arg(1))
	// 初始化服务器，返回Listener
	listener := initServer(hostAndPort)
	// 无限循环，等待客户端连接，启动协程处理
	for {
		// 等待客户端连接
		conn, err := listener.Accept()
		checkError(err, "Accept: ")
		// 启动协程处理
		go connHandler(conn)
	}
}

func initServer(hostAndPort string) net.Listener {
	// 解析tcp端点地址
	serverAddr, err := net.ResolveTCPAddr("tcp", hostAndPort)
	checkError(err, "Resolving address:port failed: '"+hostAndPort+"'")
	// 监听tcp网络
	listener, err := net.ListenTCP("tcp", serverAddr)
	checkError(err, "ListenTCP: ")
	fmt.Println("Listening to:", listener.Addr().String())
	return listener
}

func connHandler(conn net.Conn) {
	// 获取客户端地址
	connFrom := conn.RemoteAddr().String()
	fmt.Println("Connection from: ", connFrom)
	sayHello(conn)
	for {
		var ibuf []byte = make([]byte, maxRead+1)
		len, err := conn.Read(ibuf[:maxRead])
		ibuf[maxRead] = 0
		switch err {
		case nil:
			handleMsg(len, err, ibuf)
		case syscall.EAGAIN:
			// fmt.Println("run here")
			continue
			// 读取错误，断开连接
		default:
			goto DISCONNECT
		}
	}

DISCONNECT:
	err := conn.Close()
	fmt.Println("Closed connection: ", connFrom)
	checkError(err, "Close: ")
}

func sayHello(to net.Conn) {
	obuf := []byte("Let's GO!\n")
	// 发送消息给客户端
	wrote, err := to.Write(obuf)
	checkError(err, "Write: wrote "+strconv.Itoa(wrote)+" bytes.")
}

func handleMsg(len int, err error, buf []byte) {
	if len > 0 {
		fmt.Printf("<%d:", len)
		for i := 0; ; i++ {
			if buf[i] == 0 {
				fmt.Println("")
				break
			}
			fmt.Printf("%c", buf[i])
		}
	}
}

func checkError(err error, msg string) {
	if err != nil {
		panic("ERROR: " + msg + " " + err.Error())
	}
}

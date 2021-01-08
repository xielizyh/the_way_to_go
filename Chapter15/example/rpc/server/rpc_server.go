package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	rpc_objects "rpc"
	"time"
)

func main() {
	// 定义rpc_objects.Args对象
	calc := new(rpc_objects.Args)
	// 注册
	rpc.Register(calc)
	// rpc.RegisterName("Calculator", calc)
	// 调用
	rpc.HandleHTTP()
	// 监听
	listener, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Starting RPC-server -listen error:", err)
	}
	// 启动协程服务
	go http.Serve(listener, nil)
	// 延时一段时间
	time.Sleep(1000 * time.Second)
}

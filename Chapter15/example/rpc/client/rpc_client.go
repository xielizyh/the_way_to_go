package main

import (
	"fmt"
	"log"
	"net/rpc"
	rpc_objects "rpc"
)

const serverAddress = "localhost" + ":1234"

func main() {
	// 连接服务器
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		log.Fatal("Error dialing:", err)
	}
	args := &rpc_objects.Args{N: 7, M: 8}
	var reply int
	// 调用远程对象的方法
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Args error:", err)
	}
	// call1 := client.Go("Args.Multiply", args, &reply, nil)
	// <-call1.Done
	fmt.Printf("Args: %d * %d = %d", args.N, args.M, reply)
}

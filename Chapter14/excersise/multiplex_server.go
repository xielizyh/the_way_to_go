package main

import (
	"fmt"
)

type request struct {
	a, b   int
	replyc chan int
}

func (r *request) String() string {
	return fmt.Sprintf("%d+%d=%d", r.a, r.b, <-r.replyc)
}

type binOp func(a, b int) int

func run(op binOp, req *request) {
	req.replyc <- op(req.a, req.b)
}

func server(op binOp, service chan *request, quit chan bool) {
	for {
		select {
		case req := <-service:
			go run(op, req)
		case <-quit:
			return
		}
	}
}

func startServer(op binOp) (service chan *request, quit chan bool) {
	service = make(chan *request)
	quit = make(chan bool)
	go server(op, service, quit)
	return service, quit
}

func multiplexServer() {
	adder, quit := startServer(func(a, b int) int { return a + b })
	// 制作请求
	req1 := &request{3, 4, make(chan int)}
	req2 := &request{150, 250, make(chan int)}
	// 发送请求到服务端
	adder <- req1
	adder <- req2
	// 询问请求结果
	fmt.Println(req1, req2)
	quit <- true
	fmt.Println("done")
}

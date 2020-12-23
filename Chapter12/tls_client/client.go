package main

import (
	"crypto/tls"
	"fmt"
	"log"
)

func main() {
	config := &tls.Config{InsecureSkipVerify: true}

	conn, err := tls.Dial("tcp", "127.0.0.1:4430", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	fmt.Println("Handshake with server successfully established")
	n, err := conn.Write([]byte("hello server\n"))
	if err != nil {
		log.Println(err)
		return
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(buf[:n]))
}

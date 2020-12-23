package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
)

// P P
type P struct {
	X, Y, Z int
	Name    string
}

// Q Q
type Q struct {
	X, Y *int32
	Name string
}

func gob1() {
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	dec := gob.NewDecoder(&network)
	err := enc.Encode(P{3, 4, 5, "Pythoagoras"})
	if err != nil {
		log.Fatal("encode error", err)
	}
	var q Q
	err = dec.Decode(&q)
	if err != nil {
		log.Fatal("decode error", err)
	}
	// 在解码结构体的时候，只有同时匹配名称和可兼容类型的字段才会被解码
	fmt.Printf("%q: {%d, %d}\n", q.Name, *(q.X), *(q.Y))
}

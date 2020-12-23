package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

// Address 地址
type Address struct {
	Type    string
	City    string
	Country string
}

// VCard 身份证
type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func endeGob() {
	pa := &Address{"private", "Nanchong", "China"}
	wa := &Address{"work", "Chengdu", "China"}
	vc := &VCard{"Xie", "Li", []*Address{wa, pa}, "none"}
	file, _ := os.OpenFile("vcard.gob", os.O_WRONLY|os.O_CREATE, 0666)
	// defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding gob")
	}
	file.Close()

	var v VCard
	file, _ = os.OpenFile("vcard.gob", os.O_RDWR, 0600)
	// defer file.Close()
	dec := gob.NewDecoder(file)
	err = dec.Decode(&v)
	if err != nil {
		log.Println("Error in decoding gob")
	}
	fmt.Println(v)
	file.Close()
}

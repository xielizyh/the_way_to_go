package main

import (
	"encoding/json"
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

func jsonT() {
	pa := &Address{"private", "Nanchong", "China"}
	wa := &Address{"work", "Chengdu", "China"}
	vc := &VCard{"Xie", "Li", []*Address{wa, pa}, "none"}
	fmt.Printf("%v: \n", vc)
	// js, _ := json.Marshal(vc)
	js, _ := json.MarshalIndent(vc, "", "\t")
	fmt.Printf("JSON format: %s", js)
	file, _ := os.OpenFile("vcard.json", os.O_WRONLY|os.O_CREATE, 0666)
	enc := json.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
	file.Close()

	fmt.Println()
	var v1 VCard
	json.Unmarshal(js, &v1)
	fmt.Println(v1)

	var v2 VCard
	file, _ = os.OpenFile("vcard.json", os.O_RDWR, 0600)
	dec := json.NewDecoder(file)
	err = dec.Decode(&v2)
	if err != nil {
		log.Println("Error in decoding json, error:", err.Error())
	}
	fmt.Println(v2)
	file.Close()
}

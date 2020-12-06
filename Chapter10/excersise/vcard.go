package main

import (
	"fmt"
	"time"
)

// Address 地址
type Address struct {
	Street           string
	HouseNumber      uint32
	HouseNumberAddon string
	POBox            string
	ZipCode          string
	City             string
	Country          string
}

// VCard 身份证
type VCard struct {
	FirstName string
	LastName  string
	NickName  string
	BirtDate  time.Time
	Photo     string
	Addresses map[string]*Address
}

// ShowVCard 显示身份信息
func ShowVCard() {
	addr1 := &Address{"Elfenstraat", 12, "", "", "2600", "Mechelen", "België"}
	addr2 := &Address{"Heideland", 28, "", "", "2640", "Mortsel", "België"}
	addrs := make(map[string]*Address)
	addrs["youth"] = addr1
	addrs["now"] = addr2
	birtdate := time.Date(1995, 6, 3, 8, 24, 56, 0, time.Local)
	photo := "MyDocuments/MyPhotos/photo1.jpg"
	vcard := &VCard{"Ivo", "Balbaert", "", birtdate, photo, addrs}
	fmt.Printf("Here is the full VCard: %v\n", vcard)
}

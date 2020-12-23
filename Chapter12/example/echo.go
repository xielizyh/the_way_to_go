package main

import (
	"flag"
	"os"
)

// newLine 是存储Bool变量的地址
var newLine = flag.Bool("n", false, "print newline")
var verbose = flag.Bool("l", false, "list")

const (
	space   = " "
	newline = "\n"
)

func echo() {
	flag.PrintDefaults()
	flag.Parse()
	var s string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 {
			s += " "
			if *newLine {
				s += newline
			}
		}
		s += flag.Arg(i)
	}
	os.Stdout.WriteString(s)
}

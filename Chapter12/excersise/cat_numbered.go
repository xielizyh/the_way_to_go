package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

var numberflag = flag.Bool("n", false, "number each line")

func cat(r *bufio.Reader) {
	i := 0
	for {
		buf, err := r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if *numberflag {
			fmt.Fprintf(os.Stdout, "%5d %s", i, buf)
			i++
		} else {
			fmt.Fprintf(os.Stdout, "%s", buf)
		}
	}
}

func catNumbered() {
	flag.Parse()
	fmt.Printf("len(os.Args)=%d, flag.NArg()=%d\n", len(os.Args), flag.NArg())
	// flag.NArg()返回flags被处理后剩下的参数个数(flags要放在前面)
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i := 0; i < flag.NArg(); i++ {
		f, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s:error reading from %s: %s\n", os.Args[0], flag.Arg(i), err.Error())
			continue
		}
		cat(bufio.NewReader(f))
		f.Close()
	}
}

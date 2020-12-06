package main

import (
	"fmt"
	"math"
	"math/big"
	"regexp"
	"strconv"
)

func main() {
	// 9.1
	fmt.Println("------Example 9.1------")
	// 目标字符串
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	// 正则表达式
	pat := "[0-9]+.[0-9]+"

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}
	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match found!")
	}
	// 将匹配到的字符串部分替换为"##.#
	re, _ := regexp.Compile(pat)
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)
	// 参数为函数时
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
	// 9.2
	fmt.Println("------Example 9.2------")
	im := big.NewInt(math.MaxInt64)
	in := im
	io := big.NewInt(1956)
	ip := big.NewInt(1)
	ip.Mul(im, in).Add(ip, im).Div(ip, io)
	fmt.Printf("Big Int: %v\n", ip)

	// test1 := big.NewInt(1)
	test2 := big.NewInt(5)
	test3 := big.NewInt(6)
	fmt.Printf("Big Int: %v\n", test2.Add(test2, test3))

	rm := big.NewRat(math.MaxInt64, 1956)
	rn := big.NewRat(-1956, math.MaxInt64)
	ro := big.NewRat(19, 56)
	rp := big.NewRat(1111, 2222)
	rq := big.NewRat(1, 1)
	rq.Mul(rm, rn).Add(rq, ro).Mul(rq, rp)
	fmt.Printf("Big Rat: %v\n", rq)
}

package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	// 修改字符串
	str := "hello, 您好"
	strbyte := []byte(str)
	strbyte[0] = 'c'
	fmt.Println(string(strbyte))
	// 获取子串
	fmt.Println(str[:3])
	// 遍历字符串
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c ", str[i])
	}
	fmt.Println()
	for _, ch := range str {
		fmt.Printf("%c ", ch)
	}
	fmt.Println()
	// 获取字符串的字符数
	fmt.Printf("The number of str is %d\n", len([]rune(str)))
	fmt.Printf("The number of str is %d\n", utf8.RuneCountInString(str)) // 最快速
	// 连接字符串
	// https://studygolang.com/articles/12281?fr=sidebar
	var buffer bytes.Buffer // 最快速
	buffer.WriteString(str)
	buffer.WriteString(", David!")
	fmt.Println(buffer.String())
	// sl := strings.Fields(str) // 字符串分隔
	// fmt.Println("sl:", sl)
	// 第一个参数为字符串切片，第二个为字符串切片连接的分隔符
	fmt.Println(strings.Join([]string{"hello", "world"}, " "))
	// 在已有字符串数组的场合，使用 strings.Join() 能有比较好的性能
	// 在一些性能要求较高的场合，尽量使用 buffer.WriteString() 以获得更好的性能
	// 性能要求不太高的场合，直接使用运算符，代码更简短清晰，能获得比较好的可读性
	// 如果需要拼接的不仅仅是字符串，还有数字之类的其他需求的话，可以考虑 fmt.Sprintf
	os.Exit(0)
}

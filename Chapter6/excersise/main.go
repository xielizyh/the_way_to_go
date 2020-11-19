package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 6.1
	fmt.Println("------Excersise 6.1------")
	fmt.Println(Calculate1(3, 5))
	fmt.Println(Calculate1(3, 5))

	// 6.2
	fmt.Println("------Excersise 6.2------")
	ret, err := MySqrt1(0.25)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ret)
	}
	ret, err = MySqrt2(-0.25)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ret)
	}

	// 6.3
	fmt.Println("------Excersise 6.3------")
	VarArgs("hello", "world", "Thanks")
	slice := []string{"1", "2", "3"}
	// 如果参数被存储在一个 slice 类型的变量 slice 中，
	// 则可以通过 slice... 的形式来传递参数，调用变参函数
	VarArgs(slice...)
	slice1 := slice[:]
	VarArgs(slice1...)

	// 6.4
	fmt.Println("------Excersise 6.4------")
	pos, res := Fibonacci(10)
	fmt.Printf("fibonacci(%d) is: %d\n", pos, res)

	// 6.5
	fmt.Println("------Excersise 6.5------")
	RecursionPrint(10)
	fmt.Println("")

	// 6.6
	fmt.Println("------Excersise 6.6------")
	fmt.Printf("Factorial(5) is %d\n", Factorial(5))

	// 6.7
	fmt.Println("------Excersise 6.7------")
	spaceIdx := strings.IndexFunc("hello world", unicode.IsSpace)
	fmt.Printf("The first space index is: %d\n", spaceIdx)
	str := "hello 你好 world"
	fmt.Printf("The origin string is %s\n", str)
	newStr := RelpaceNotASCII(str)
	fmt.Printf("After replace not ASCII, the str is: %s\n", newStr)

	// 6.8
	fmt.Println("------Excersise 6.8------")
	fv := func() (ret int) {
		fmt.Printf("the origin ret is: %d\n", ret)
		defer func() {
			ret++
		}()
		fmt.Println("Hello World")
		return 1
	}
	value := fv()
	fmt.Printf("fv is of type %T, value is %d\n", fv, value)
	// 6.9
	fmt.Println("------Excersise 6.9------")
	fv1 := Fibonacci1()
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", fv1(i))
	}
	fmt.Println()
	// 6.10
	fmt.Println("------Excersise 6.10------")
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpg")
	fmt.Println(addBmp("file"), addJpeg("file"), addBmp("test.bmp"))

}

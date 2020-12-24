package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ParseError 解析失败
type ParseError struct {
	Index int
	Word  string
	Err   error
}

// String 显示
func (e *ParseError) String() string {
	return fmt.Sprintf("pkg parse: error parsing %q as integer", e.Word)
}

func fields2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no words to parse")
	}
	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(&ParseError{idx, field, err})
		}
		numbers = append(numbers, num)
	}
	return
}

// Parse 解析
func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			if err, ok = r.(error); !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()

	fields := strings.Fields(input)
	numbers = fields2numbers(fields)
	return
}

// 将 panic 转换成 error 来告诉调用方为何出错，是很实用的
func panicPackage() {
	var examples = []string{
		"1 2 3 4 5",
		"100 50 25 12.5 6.25",
		"2 + 2 = 4",
		"1st class",
		"",
	}
	for _, ex := range examples {
		fmt.Printf("Parsing %q:\n ", ex)
		nums, err := Parse(ex)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(nums)
	}
}

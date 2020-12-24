package main

import (
	"fmt"
	"math"
)

// ConvertInt64ToInt 转换成int
func ConvertInt64ToInt(num int64) int {
	if math.MinInt32 < num && num <= math.MaxInt32 {
		return int(num)
	}
	panic(fmt.Sprintf("%d is out of the int32 range", num))
}

// IntFromInt64 转换
func IntFromInt64(num int64) (n int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("%v", r)
		}
	}()
	n = ConvertInt64ToInt(num)
	return n, nil
}

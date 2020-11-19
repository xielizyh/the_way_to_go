package main

import (
	"errors"
	"math"
)

// MySqrt1 返回平方根
func MySqrt1(num float64) (float64, error) {
	if num < 0 {
		return float64(math.NaN()), errors.New("I won't be able to do a sqrt of negative number")
	}

	return math.Sqrt(num), nil
}

// MySqrt2 返回平方根
func MySqrt2(num float64) (result float64, err error) {
	if num < 0 {
		result = float64(math.NaN())
		err = errors.New("I won't be able to do a sqrt of negative number")
		return
	}

	result = math.Sqrt(num)
	return
}

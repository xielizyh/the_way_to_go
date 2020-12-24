package main

import "fmt"

func recoverDividebyZero(dividend, divisor int) (quotient float64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panicing %s", r)
		}
	}()
	quotient = float64(dividend / divisor)
	return
}

package main

import (
	"errors"
)

// RemoveStringSlice 从切片中移除元素
func RemoveStringSlice(slice []int, start, end uint16) (ret []int, err error) {
	if start > end || int(end) > len(slice) {
		err = errors.New("the index of start or end is out of range")
		return
	}
	n := len(slice) - int(end-start+1)
	copy(slice[start:], slice[end+1:])

	return slice[:n], nil
}

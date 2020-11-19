package main

import (
	"strings"
)

func isNotASCII(r rune) bool {
	if r > 255 {
		return true
	}
	return false
}

// RelpaceNotASCII 替换非ASCII字符为?
func RelpaceNotASCII(s string) string {
	newStr := s[:]
	idx := 0
	for {
		idx = strings.IndexFunc(newStr, isNotASCII)
		if idx == -1 {
			break
		}
		newStr = newStr[:idx] + "?" + newStr[idx+1:]
	}

	return newStr
}

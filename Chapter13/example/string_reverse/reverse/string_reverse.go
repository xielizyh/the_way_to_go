package reverse

// StringReverse 字符串反序
func StringReverse(s string) string {
	// 转化为utf-8
	runes := []rune(s)
	n, h := len(runes), len(runes)/2
	for i := 0; i < h; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

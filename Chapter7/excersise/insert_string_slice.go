package main

// InsertStringSlice 将切片插入到另一个切片的指定位置
func InsertStringSlice(src []string, dst []string, offset int) []string {
	m := len(src)
	n := m + len(dst)
	tmp := src[offset:]
	if n > cap(src) {
		newSlice := make([]string, (n+1)*2)
		copy(newSlice, src[:offset])
		src = newSlice
	}
	copy(src[offset:], dst[:])
	copy(src[offset+len(dst):], tmp)

	return src[0:n]
}

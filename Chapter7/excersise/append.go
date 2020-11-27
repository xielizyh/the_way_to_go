package main

// Append 追加
func Append(slice, data []byte) []byte {
	slice = append(slice, data...)
	// fmt.Println(slice)
	return slice
}

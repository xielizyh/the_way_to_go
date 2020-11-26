package main

// ArrayValue 数组内存拷贝
func ArrayValue(arr [3]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
}

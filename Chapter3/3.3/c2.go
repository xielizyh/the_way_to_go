package main

// #include <stdlib.h>
// #include <stdio.h>
import "C"
import "unsafe"

func print(s string) {
	// 将go中string转换成c
	cs := C.CString(s)
	// 手动调用 C.free 来释放变量的内存
	// Go 的内存管理机制无法管理通过 C 代码分配的内存
	defer C.free(unsafe.Pointer(cs))
	C.fputs(cs, (*C.FILE)(C.stdout))
}

func main() {
	print("Go调用C中string")
}

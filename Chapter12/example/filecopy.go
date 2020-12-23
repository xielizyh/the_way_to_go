package main

import (
	"io"
	"os"
)

// CopyFile 拷贝文件
func CopyFile(dstFile string, srcFile string) (written int64, err error) {
	src, err := os.Open(srcFile)
	if err != nil {
		return -1, err
	}
	defer src.Close()
	dst, err := os.OpenFile(dstFile, os.O_WRONLY, 0666)
	if err != nil {
		return -1, err
	}
	defer dst.Close()
	return io.Copy(dst, src)
}

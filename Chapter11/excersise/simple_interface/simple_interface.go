package simpleinterface

import "fmt"

// Simpler 接口
type Simpler interface {
	Get() int
	Set(n int)
}

// Simple 类型
type Simple struct {
	Num int
}

// Get 获取
func (s *Simple) Get() int {
	return s.Num
}

// Set 设置
func (s *Simple) Set(n int) {
	s.Num = n
}

// SimpleFunc 测试
func SimpleFunc(ser Simpler) {
	ser.Set(5)
	fmt.Println(ser.Get())
}

// RSimple RSimple
type RSimple struct{}

// Get 获取
func (r *RSimple) Get() int {
	return -1
}

// Set 设置
func (r *RSimple) Set(n int) {

}

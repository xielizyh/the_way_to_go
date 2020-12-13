package floatsort

import (
	"fmt"
	"math/rand"
	"time"
)

// Sorter 排序接口
type Sorter interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// Sort 排序
func Sort(data Sorter) {
	for pass := 1; pass <= data.Len(); pass++ {
		for i := 0; i < data.Len()-pass; i++ {
			if data.Less(i+1, i) {
				data.Swap(i, i+1)
			}
		}
	}
}

// IsSorted 是否排序
func IsSorted(data Sorter) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}

//Float64Array float64切片
type Float64Array []float64

// Len 个数
// 由于Float64Array是切片类型（引用类型），因此不用像结构体那样为了不"拷贝"，设置接收者类型为指针接收
// func (p *Float64Array) Len() int { return len(*p) }
func (p Float64Array) Len() int { return len(p) }

// Less 比较大小
func (p Float64Array) Less(i, j int) bool { return p[i] < p[j] }

// Swap 交换
func (p Float64Array) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// NewFloat64Array 创建数组
func NewFloat64Array(num int) Float64Array {
	array := make(Float64Array, num)

	return array
}

// List 打印数组
func (p Float64Array) List() string {
	s := "{ "
	for i := 0; i < p.Len(); i++ {
		s += fmt.Sprintf("%3.1f ", p[i])
	}
	s += " }"

	return s
}

// String 字符串
func (p Float64Array) String() string {
	return p.List()
}

// Fill 创建
func (p Float64Array) Fill(num int) {
	timens := int64(time.Now().Nanosecond())
	rand.Seed(timens)
	for i := 0; i < num; i++ {
		p[i] = 100 * rand.Float64()
	}
}

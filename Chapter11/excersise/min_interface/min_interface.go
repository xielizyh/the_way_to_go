package mininterface

// Miner 最小值接口
type Miner interface {
	Len() int
	ElemIx(ix int) interface{}
	Less(i, j int) bool
	Swap(i, j int)
}

// Min 最小值
func Min(data Miner) interface{} {
	min := data.ElemIx(0)
	for i := 1; i < data.Len(); i++ {
		if data.Less(i, i-1) {
			min = data.ElemIx(i)
		} else {
			data.Swap(i, i-1)
		}
	}
	return min
}

// IntArray 整数切片
type IntArray []int

// Len 长度
func (p IntArray) Len() int { return len(p) }

// ElemIx 元素
func (p IntArray) ElemIx(ix int) interface{} { return p[ix] }

// Less 比较
func (p IntArray) Less(i, j int) bool { return p[i] < p[j] }

// Swap 交换
func (p IntArray) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

// StringArray 整数切片
type StringArray []string

// Len 长度
func (p StringArray) Len() int { return len(p) }

// ElemIx 元素
func (p StringArray) ElemIx(ix int) interface{} { return p[ix] }

// Less 比较
func (p StringArray) Less(i, j int) bool { return p[i] < p[j] }

// Swap 交换
func (p StringArray) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

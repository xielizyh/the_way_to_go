package stackgeneral

import (
	"errors"
	"fmt"
)

// Stack 堆栈
type Stack []interface{}

// Len 堆栈深度
func (s Stack) Len() int { return len(s) }

// Cap 堆栈容量
func (s Stack) Cap() int { return cap(s) }

// IsEmpty 堆栈是否空
func (s Stack) IsEmpty() bool { return len(s) == 0 }

// PushTmp 入栈
func (s Stack) PushTmp(e interface{}) {
	fmt.Printf("the s1 is %p\n", s)
	s = append(s, e)
	fmt.Printf("the s3 is %p\n", s)
}

// Push 入栈
// 注意：为什么用指针接收值？因为用值接收时，因为Push方法中调用了append函数，
// append会在原始数组容量不足的情况下重新分配空间, Go 默认会先开一片内存区域，
// 把原来的值拷贝过来，然后再执行 append() 操作。这种情况丝毫不影响原数组。
// 原始切片的指针，长度，容量却没有改变。
func (s *Stack) Push(e interface{}) {
	*s = append(*s, e)
}

// Top 出栈
func (s Stack) Top() (interface{}, error) {
	if len(s) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s[len(s)-1], nil
}

// Pop 出栈
func (s *Stack) Pop() (interface{}, error) {
	stk := *s
	if len(stk) == 0 {
		return nil, errors.New("stack is empty")
	}
	top := stk[len(stk)-1]
	*s = stk[:len(stk)-1]
	return top, nil
}

// String stack方法
func (s Stack) String() string {
	ret := ""
	for i := 0; i < len(s); i++ {
		switch s[i].(type) {
		case int:
			ret += fmt.Sprintf("[%d:%d] ", i, s[i].(int))
		case string:
			ret += fmt.Sprintf("[%d:%s] ", i, s[i].(string))
		case float64:
			ret += fmt.Sprintf("[%d:%.2f] ", i, s[i].(float64))
		}
	}
	return ret
}

/* // LIMIT 栈大小
const LIMIT = 4

// Stack 结构体实现
type Stack struct {
	data [LIMIT]interface{}
	sp   int
}

// String stack方法
func (s Stack) String() string {
	ret := ""
	for i := 0; i < s.sp; i++ {
		switch s.data[i].(type) {
		case int:
			ret += fmt.Sprintf("[%d:%d] ", i, s.data[i].(int))
		case string:
			ret += fmt.Sprintf("[%d:%s] ", i, s.data[i].(string))
		case float64:
			ret += fmt.Sprintf("[%d:%.2f] ", i, s.data[i].(float64))
		}
	}
	return ret
}

// Pop 出栈
func (s *Stack) Pop() interface{} {
	if s.sp < 0 {
		return -1
	}
	defer func() {
		s.data[s.sp] = 0
	}()
	s.sp--
	return s.data[s.sp]
}

// Push 入栈
func (s *Stack) Push(n interface{}) int {
	if s.sp >= LIMIT {
		return -1
	}
	defer func() {
		s.sp++
	}()
	s.data[s.sp] = n
	return 0
}
*/

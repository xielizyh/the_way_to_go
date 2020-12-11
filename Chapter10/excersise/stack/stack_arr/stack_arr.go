package stackarr

import "fmt"

// LIMIT 栈大小
const LIMIT = 4

// Stack 数组作为底层结构
type Stack [LIMIT]int

var sp int = 0

// String stack方法
func (s Stack) String() string {
	ret := ""
	for i := 0; i < sp; i++ {
		ret += fmt.Sprintf("[%d:%d] ", i, s[i])
	}
	return ret
}

// Pop 出栈
func (s *Stack) Pop() int {
	if sp < 0 {
		return -1
	}
	defer func() {
		s[sp] = 0
	}()
	sp--
	return s[sp]
}

// Push 入栈
// 如果想要方法改变接收者的数据，就在接收者的指针类型上定义该方法.
func (s *Stack) Push(n int) int {
	if sp >= len(*s) {
		return -1
	}
	defer func() {
		sp++
	}()
	s[sp] = n
	return 0
}

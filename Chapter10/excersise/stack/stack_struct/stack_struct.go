package stackstruct

import "fmt"

// LIMIT 栈大小
const LIMIT = 4

// Stack 结构体实现
type Stack struct {
	data [LIMIT]int
	sp   int
}

// String stack方法
func (s Stack) String() string {
	ret := ""
	for i := 0; i < s.sp; i++ {
		ret += fmt.Sprintf("[%d:%d] ", i, s.data[i])
	}
	return ret
}

// Pop 出栈
func (s *Stack) Pop() int {
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
func (s *Stack) Push(n int) int {
	if s.sp >= LIMIT {
		return -1
	}
	defer func() {
		s.sp++
	}()
	s.data[s.sp] = n
	return 0
}

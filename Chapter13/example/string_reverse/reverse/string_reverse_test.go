package reverse

import "testing"

type ReverseTest struct {
	in, out string
}

var ReverseTests = []ReverseTest{
	{"Hello World", "dlroW olleH"},
	{"您好！小明", "明小！好您"},
	{"Hello 世界", "界世 olleH"},
}

// 单元测试的重点在于发现程序设计或实现的逻辑错误
// https://studygolang.com/articles/7051
func TestStringReverse(t *testing.T) {
	/* 	if StringReverse("Hello World") != "dlroW olleH" {
	   		t.Log("After reverse, \"Hello World\" must be \"dlroW olleH\"")
	   		t.Fail()
	   	}
	   	if StringReverse("您好！小明") != "明小！好您" {
	   		t.Log("After reverse, \"您好！小明\" must be \"明小！好您\"")
	   		t.Fail()
	   	} */
	for _, r := range ReverseTests {
		exp := StringReverse(r.in)
		if r.out != exp {
			t.Errorf("Reverse of %s expects %s, but got %s", r.in, exp, r.out)
		}
	}
}

// 基准测试(压力测试)用来检测函数(方法）的性能
// 执行方式 go test -v -bench=.
// 代码中的函数会被调用 N 次（N是非常大的数，如 N = 1000000），
// 并展示 N 的值和函数执行的平均时间，单位为 ns（纳秒，ns/op）
func BenchmarkStringReverse(b *testing.B) {
	s := "ABCD"
	for i := 0; i < b.N; i++ {
		StringReverse(s)
	}
}

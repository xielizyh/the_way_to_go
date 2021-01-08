package key

import (
	"testing"
)

type GenKeyTest struct {
	n   int
	in  string
	out []byte
}

var GenKeyTests = []GenKeyTest{
	{0, "hello", []byte{0x22, 0x7d, 0xa5, 0x38, 0xae, 0x11, 0xe1, 0x7b, 0x60, 0x75, 0xfe, 0x59, 0xb2, 0x02, 0x5f, 0x7e}},
	{65536, "Test", []byte{0x21, 0x4d, 0xef, 0x9c, 0x54, 0xbd, 0x9a, 0xbf, 0x0c, 0x28, 0xf5, 0xab, 0x48, 0x2e, 0x76, 0xe4}},
}

func TestGetKey(t *testing.T) {
	for _, val := range GenKeyTests {
		exp := GenKey(val.n, val.in)
		if len(exp) == len(val.out) {
			for i, v := range exp {
				if v != val.out[i] {
					t.Errorf("expects %v, but got %v", val.out, exp)
				}
			}
		} else {
			t.Errorf("expects len %d, but got %v", len(val.out), len(exp))
		}
	}
}

func BenchmarkGenKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenKey(GenKeyTests[0].n, GenKeyTests[0].in)
	}
}

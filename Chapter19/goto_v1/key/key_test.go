package key

import (
	"testing"
)

type GenKeyTest struct {
	n       int
	in, out string
}

var GenKeyTests = []GenKeyTest{
	{0, "hello", "world"},
	{65536, "Test", "Test"},
}

func TestGetKey(t *testing.T) {
	for _, val := range GenKeyTests {
		exp := GenKey(val.n, val.in)
		if val.out != exp {
			t.Errorf("got %s, want %s", exp, val.out)
		}
	}
}

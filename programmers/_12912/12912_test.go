package _12912

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		a      int
		b      int
		expect int64
	}{
		{3, 5, 12},
		{3, 3, 3},
		{5, 3, 12},
	}
	for _, c := range cases {
		if r := solution(c.a, c.b); r != c.expect {
			t.Errorf("a %#v, b %#v; got %#v, want %#v", c.a, c.b, r, c.expect)
		}
	}
}

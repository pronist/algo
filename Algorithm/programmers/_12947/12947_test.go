package _12947

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		n      int
		expect bool
	}{
		{10, true},
		{12, true},
		{11, false},
	}
	for _, c := range cases {
		if r := solution(c.n); r != c.expect {
			t.Errorf("n %#v; got %#v, want %#v", c.n, r, c.expect)
		}
	}
}
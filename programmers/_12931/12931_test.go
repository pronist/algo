package _12931

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		n      int
		expect int
	}{
		{123, 6},
		{987, 24},
	}
	for _, c := range cases {
		if r := solution(c.n); r != c.expect {
			t.Errorf("n %#v; got %#v, want %#v", c.n, r, c.expect)
		}
	}
}

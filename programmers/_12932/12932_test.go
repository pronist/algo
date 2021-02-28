package _12932

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct{
		n int64
		expect []int
	} {
		{12345, []int{5,4,3,2,1}},
		{123456, []int{6,5,4,3,2,1}},
	}
	for _, c := range cases {
		r := solution(c.n)
		if len(r) != len(c.expect) {
			t.Errorf("got %#v, want %#v", r, c.expect)
			return
		}
		for i, v := range r {
			if v != c.expect[i] {
				t.Errorf("n %#v; got %#v, want %#v", c.n, v, c.expect[i])
			}
		}
	}
}
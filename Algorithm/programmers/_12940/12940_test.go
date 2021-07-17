package _12940

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct{
		n int
		m int
		expect []int
	} {
		{3, 12, []int{3, 12}},
		{2, 5, []int{1, 10}},
		{10, 20, []int{10, 20}},
		{1, 1, []int{1, 1}},
	}
	for _, c := range cases {
		r := solution(c.n, c.m);
		if len(r) != len(c.expect) {
			t.Errorf("got %#v, want %#v", r, c.expect)
			return
		}
		for i, v := range r {
			if v != c.expect[i] {
				t.Errorf("m %#v, n %#v; got %#v, want %#v", c.m, c.n, r, c.expect)
			}
		}
	}
}

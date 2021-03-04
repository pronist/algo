package _12954

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		x      int
		n      int
		expect []int64
	}{
		{2, 5, []int64{2, 4, 6, 8, 10}},
		{4, 3, []int64{4, 8, 12}},
		{-4, 2, []int64{-4, -8}},
	}
	for _, c := range cases {
		r := solution(c.x, c.n)
		if len(r) != len(c.expect) {
			t.Errorf("got %#v, want %#v", r, c.expect)
			return
		}
		for i, v := range r {
			if v != c.expect[i] {
				t.Errorf("x %#v, n %#v; got %#v, want %#v", c.x, c.n, r, c.expect)
			}
		}
	}
}

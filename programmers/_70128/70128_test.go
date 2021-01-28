package _70128

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		a      []int
		b      []int
		expect int
	}{
		{[]int{1, 2, 3, 4}, []int{-3, -1, 0, 2}, 3},
		{[]int{-1, 0, 1}, []int{1, 0, -1}, -2},
	}
	for _, c := range cases {
		if r := solution(c.a, c.b); r != c.expect {
			t.Errorf("a %#v, b %#v; got %#v, want %#v", c.a, c.b, r, c.expect)
		}
	}
}

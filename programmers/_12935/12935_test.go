package _12935

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		arr    []int
		expect []int
	}{
		{[]int{4, 3, 2, 1}, []int{4, 3, 2}},
		{[]int{10}, []int{-1}},
	}
	for _, c := range cases {
		r := solution(c.arr)
		if len(r) != len(c.expect) {
			t.Errorf("got %#v, want %#v", r, c.expect)
			return
		}
		for i, v := range r {
			if v != c.expect[i] {
				t.Errorf("arr %#v; got %#v, want %#v", c.arr, v, c.expect[i])
			}
		}
	}
}

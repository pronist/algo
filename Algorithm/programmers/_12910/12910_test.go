package _12910

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		arr     []int
		divisor int
		expect  []int
	}{
		{[]int{5, 9, 7, 10}, 5, []int{5, 10}},
		{[]int{2, 36, 1, 3}, 1, []int{1, 2, 3, 36}},
		{[]int{3, 2, 6}, 5, []int{-1}},
	}
	for _, c := range cases {
		r := solution(c.arr, c.divisor)
		if len(r) != len(c.expect) {
			t.Errorf("got %#v, want %#v", r, c.expect)
			return
		}
		for i, v := range r {
			if v != c.expect[i] {
				t.Errorf("#%d, got %#v, want %#v", i, r, c.expect)
			}
		}
	}
}

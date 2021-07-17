package _42862

import "testing"

func TestRemove(t *testing.T) {
	cases := []struct {
		s      []int
		i      int
		expect []int
	}{
		{[]int{0, 1, 2}, 0, []int{1, 2}},
		{[]int{0, 1, 2}, 1, []int{0, 2}},
		{[]int{0, 1, 2}, 2, []int{0, 1}},
	}
	for _, c := range cases {
		c.s = remove(c.s, c.i)
		if len(c.s) != len(c.expect) {
			t.Errorf("got %#v, want %#v", c.s, c.expect)
			return
		}
		for i, v := range c.s {
			if v != c.s[i] {
				t.Errorf("#%d, got %#v, want %#v", i, c.s, c.expect)
			}
		}
	}
}

func TestSolution(t *testing.T) {
	cases := []struct {
		n       int
		lost    []int
		reserve []int
		expect  int
	}{
		{5, []int{2, 4}, []int{1, 3, 5}, 5},
		{5, []int{2, 4}, []int{3}, 4},
		{3, []int{3}, []int{1}, 2},
		{5, []int{2, 4, 5}, []int{4, 1}, 4},
		{5, []int{2, 4}, []int{2, 3, 5}, 5},
		{5, []int{2, 3, 5}, []int{2, 3}, 4},
		{5, []int{1, 2, 4, 5}, []int{2, 3, 4}, 3},
	}
	for _, c := range cases {
		if r := solution(c.n, c.lost, c.reserve); r != c.expect {
			t.Errorf("n: %#v, lost: %#v, reserve: %#v; got %#v, want %#v", c.n, c.lost, c.reserve, r, c.expect)
		}
	}
}

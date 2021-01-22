package _42748

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		array    []int
		commands [][]int
		expect   []int
	}{
		{[]int{1, 5, 2, 6, 3, 7, 4}, [][]int{{2, 5, 3}, {4, 4, 1}, {1, 7, 3}}, []int{5, 6, 3}},
	}
	for _, c := range cases {
		r := solution(c.array, c.commands)
		if len(r) != len(c.expect) {
			t.Errorf("got %#v, want %#v", r, c.expect)
			return
		}
		for i, v := range r {
			if v != c.expect[i] {
				t.Errorf("#%d, got %#v, want %#v", i, r, c.expect)
				return
			}
		}
	}
}

package _42587

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		priorities []int
		location   int
		expect     int
	}{
		{[]int{2, 1, 3, 2}, 2, 1},
		{[]int{1, 1, 9, 1, 1, 1}, 0, 5},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 4, 3},
		{[]int{1, 1, 1, 1}, 3, 4},
		{[]int{1, 1, 1, 2}, 1, 3},
		{[]int{2, 2, 2, 1, 3, 4}, 3, 6},
	}
	for _, c := range cases {
		if r := solution(c.priorities, c.location); r != c.expect {
			t.Errorf("priorities %#v; got %#v, want %#v", c.priorities, r, c.expect)
		}
	}
}

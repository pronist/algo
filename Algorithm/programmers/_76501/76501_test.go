package _76501

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		absolutes []int
		signs     []bool
		expect    int
	}{
		{[]int{4, 7, 12}, []bool{true, false, true}, 9},
		{[]int{1, 2, 3}, []bool{false, false, true}, 0},
	}

	for _, c := range cases {
		if r := solution(c.absolutes, c.signs); r != c.expect {
			t.Errorf("absolutes %#v, signs %#v; got %#v, want %#v", c.absolutes, c.signs, r, c.expect)
		}
	}
}

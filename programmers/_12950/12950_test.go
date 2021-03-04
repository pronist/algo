package _12950

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		arr1   [][]int
		arr2   [][]int
		expect [][]int
	}{
		{
			[][]int{{1, 2}, {2, 3}},
			[][]int{{3, 4}, {5, 6}},
			[][]int{{4, 6}, {7, 9}},
		},
		{
			[][]int{{1}, {2}},
			[][]int{{3}, {5}},
			[][]int{{4}, {7}},
		},
	}
	for _, c := range cases {
		r := solution(c.arr1, c.arr2)
		if len(r) != len(c.expect) {
			t.Errorf("got %#v, want %#v", r, c.expect)
			return
		}
		for i := 0; i < len(c.expect); i++ {
			for j := 0; j < len(c.expect[i]); j++ {
				if r[i][j] != c.expect[i][j] {
					t.Errorf("arr1 %#v, arr2 %#v; got %#v, want %#v", c.arr1, c.arr2, r, c.expect)
				}
			}
		}
	}
}

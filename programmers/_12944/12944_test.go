package _12944

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		arr    []int
		expect float64
	}{
		{[]int{1,2,3,4}, 2.5},
		{[]int{5,5}, 5},
	}
	for _, c := range cases {
		if r := solution(c.arr); r != c.expect {
			t.Errorf("n %#v; got %#v, want %#v", c.arr, r, c.expect)
		}
	}
}

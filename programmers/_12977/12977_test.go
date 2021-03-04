package _12977

import "testing"

func TestIsPrime(t *testing.T) {
	cases := []struct {
		n      int
		expect bool
	}{
		{2, true},
		{5, true},
		{10, false},
		{11, true},
		{13, true},
		{15, false},
	}

	for _, c := range cases {
		if r := isPrime(c.n); r != c.expect {
			t.Errorf("n %#v; got %#v, want %#v", c.n, r, c.expect)
		}
	}
}

func TestSolution(t *testing.T) {
	cases := []struct {
		nums   []int
		expect int
	}{
		{[]int{1, 2, 3, 4}, 1},
		{[]int{1, 2, 7, 6, 4}, 4},
	}
	for _, c := range cases {
		if r := solution(c.nums); r != c.expect {
			t.Errorf("nums %#v; got %#v, want %#v", c.nums, r, c.expect)
		}
	}
}

package _12943

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		n      int
		expect int
	}{
		{6, 8},
		{16, 4},
		{626331, -1},
	}
	for _, c := range cases {
		if r := solution(c.n); r != c.expect {
			t.Errorf("n %#v; got %#v, want %#v", c.n, r, c.expect)
		}
	}
}
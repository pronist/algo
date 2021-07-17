package _12934

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		n      int64
		expect int64
	}{
		{121, 144},
		{3, -1},
	}
	for _, c := range cases {
		if r := solution(c.n); r != c.expect {
			t.Errorf("n %#v; got %#v, want %#v", c.n, r, c.expect)
		}
	}
}

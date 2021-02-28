package _12933

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		n      int64
		expect int64
	}{
		{118372, 873211},
	}
	for _, c := range cases {
		if r := solution(c.n); r != c.expect {
			t.Errorf("n %#v; got %#v, want %#v", c.n, r, c.expect)
		}
	}
}

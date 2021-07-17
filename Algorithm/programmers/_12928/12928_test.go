package _12928

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct{
		n int
		expect int
	} {
		{12, 28},
		{5, 6},
	}
	for _, c := range cases {
		if r := solution(c.n); r != c.expect {
			t.Errorf("n %#v; got %#v, want %#v", c.n, r, c.expect)
		}
	}
}

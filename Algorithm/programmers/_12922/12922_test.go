package _12922

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct{
		n int
		expect string
	} {
		{3, "수박수"},
		{4, "수박수박"},
	}
	for _, c := range cases {
		if r := solution(c.n); r != c.expect {
			t.Errorf("n %#v; got %#v, want %#v", c.n, r, c.expect)
		}
	}
}

package _12926

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct{
		s string
		n int
		expect string
	} {
		{"AB", 1, "BC"},
		{"z", 1, "a"},
		{"a B z", 4, "e F d"},
		{"Z", 2, "B"},
		{"z", 3, "c"},
		{"y", 2, "a"},
	}
	for _, c := range cases {
		if r := solution(c.s, c.n); r != c.expect {
			t.Errorf("s %#v, n %#v; got %#v, want %#v", c.s, c.n, r, c.expect)
		}
	}
}

package _12925

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct{
		s string
		expect int
	} {
		{"1234", 1234},
		{"-1234", -1234},
	}
	for _, c := range cases {
		if r := solution(c.s); r != c.expect {
			t.Errorf("s %#v; got %#v, want %#v", c.s, r, c.expect)
		}
	}
}

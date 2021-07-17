package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		a      int
		b      int
		expect string
	}{
		{2, 2, "**\n**\n"},
		{5, 3, "*****\n*****\n*****\n"},
	}
	for _, c := range cases {
		buf := solution(c.a, c.b)
		if buf.String() != c.expect {
			t.Errorf("got %#v, want %#v", buf.String(), c.expect)
		}
	}
}

package _12918

import (
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		s      string
		expect bool
	} {
		{"a234", false},
		{"1234", true},
	}
	for _, c := range cases {
		if r := solution(c.s); r != c.expect {
			t.Errorf("s %#v; got %#v, want %#v", c.s, r, c.expect)
		}
	}
}
package _12930

import (
	"testing"
)

func TestSolution(t *testing.T) {
	cases := []struct {
		s      string
		expect string
	}{
		{"try hello world", "TrY HeLlO WoRlD"},
		{"try", "TrY"},
	}
	for _, c := range cases {
		if r := solution(c.s); r != c.expect {
			t.Errorf("s %#v; got %#v, want %#v", c.s, r, c.expect)
		}
	}
}

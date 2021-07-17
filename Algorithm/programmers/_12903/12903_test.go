package _12903

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		s      string
		expect string
	}{
		{"abcde", "c"},
		{"qwer", "we"},
	}
	for _, c := range cases {
		if r := solution(c.s); r != c.expect {
			t.Errorf("s %#v; got %#v, want %#v", c.s, r, c.expect)
		}
	}
}

package _12917

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		s      string
		expect string
	}{
		{"Zbcdefg", "gfedcbZ"},
		{"ZbcdWRWWefag", "gfedcbaZWWWR"},
	}
	for _, c := range cases {
		if r := solution(c.s); r != c.expect {
			t.Errorf("got %#v, want %#v", r, c.expect)
			return
		}
	}
}

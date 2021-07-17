package _12915

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		strings []string
		n       int
		expect  []string
	}{
		{[]string{"sun", "bed", "car"}, 1, []string{"car", "bed", "sun"}},
		{[]string{"abce", "abcd", "cdx"}, 2, []string{"abcd", "abce", "cdx"}},
	}
	for _, c := range cases {
		r := solution(c.strings, c.n)
		if len(r) != len(c.expect) {
			t.Errorf("got %#v, want %#v", r, c.expect)
			return
		}
		for i, v := range r {
			if v != c.expect[i] {
				t.Errorf("got %#v, want %#v", v, c.expect[i])
			}
		}
	}
}

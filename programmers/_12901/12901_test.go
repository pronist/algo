package _12901

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		month  int
		day    int
		expect string
	}{
		{5, 24, "TUE"},
	}
	for _, c := range cases {
		if r := solution(c.month, c.day); r != c.expect {
			t.Errorf("Month %#v, Day %#v; got %#v, want %#v", c.month, c.day, r, c.expect)
		}
	}
}

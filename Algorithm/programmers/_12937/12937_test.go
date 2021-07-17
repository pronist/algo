package _12937

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		num    int
		expect string
	}{
		{3, "Odd"},
		{4, "Even"},
	}
	for _, c := range cases {
		if r := solution(c.num); r != c.expect {
			t.Errorf("num %#v; got %#v, want %#v", c.num, r, c.expect)
		}
	}
}

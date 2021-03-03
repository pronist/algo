package _12948

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		n      string
		expect string
	}{
		{"01033334444", "*******4444"},
		{"027778888", "*****8888"},
	}
	for _, c := range cases {
		if r := solution(c.n); r != c.expect {
			t.Errorf("n %#v; got %#v, want %#v", c.n, r, c.expect)
		}
	}
}
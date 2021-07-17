package _12919

import "testing"

func TestSolution(t *testing.T) {
	cases := []struct {
		seoul  []string
		expect string
	}{
		{[]string{"Jane", "Kim"}, "김서방은 1에 있다"},
	}
	for _, c := range cases {
		if r := solution(c.seoul); r != c.expect {
			t.Errorf("seoul %#v; got %#v, want %#v", c.seoul, r, c.expect)
		}
	}
}

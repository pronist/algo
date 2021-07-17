package main

import (
	"testing"
)

func TestSub(t *testing.T) {
	cases := []struct{
		a int
		b int
		expect int
	} {
		{3, 2, 1},
	}
	for _, c := range cases {
		if r := sub(c.a, c.b); r != c.expect {
			t.Errorf("a %#v, b %#v; got %#v, want %#v", c.a, c.b, r, c.expect)
		}
	}
}

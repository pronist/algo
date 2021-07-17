package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	cases := []struct{
		a int
		b int
		expect int
	} {
		{1, 2, 3},
	}
	for _, c := range cases {
		if r := add(c.a, c.b); r != c.expect {
			t.Errorf("a %#v, b %#v; got %#v, want %#v", c.a, c.b, r, c.expect)
		}
	}
}

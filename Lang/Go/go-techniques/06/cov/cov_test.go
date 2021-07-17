package cov

import (
	"strings"
	"testing"
)

type Case struct {
	in, out string
}

var Cases = []Case{
	{"", "wordless?"},
	{"날", "one word"},
	{"Go!", "a few words"},
	{"Go Go!", "many words"},
	{strings.Repeat("*", 10), "too many words"},
}

func TestWords(t *testing.T) {
	for i, c := range Cases {
		r := Words(c.in)
		if r != c.out {
			t.Errorf("#%d: Words(%s) got %s; want %s", i, c.in, r, c.out)
		}
	}
}
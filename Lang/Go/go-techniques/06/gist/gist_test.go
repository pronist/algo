package gist

import (
	"io"
	"strings"
	"testing"
)

type dummyDoer struct {}

func (d *dummyDoer) doGistRequest(user string) (io.Reader, error) {
	return strings.NewReader(`
[
	{"html_url": "https://gist.github.com/examples.1"},
	{"html_url": "https://gist.github.com/examples.2"}
]`), nil
}

func TestListGists(t *testing.T) {
	c := &Client{&dummyDoer{}}
	urls, err := c.ListGists("test")
	if err != nil {
		t.Fatal(err)
	}
	if expected := 2; len(urls) != expected {
		t.Fatalf("want %d, got %d", expected, len(urls))
	}
}
package hello

import "testing"

func TestSayHello(t *testing.T) {
	if s := SayHello(); s != "Hello, Go!" {
		t.Errorf("SayHello() got %s; want %s", "Hello, Go!", s)
	}
}

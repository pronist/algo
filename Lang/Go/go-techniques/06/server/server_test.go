package server

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoute(t *testing.T) {
	ts := httptest.NewServer(Route())
	defer ts.Close()

	res, err := http.Get(ts.URL + "/greet?name=gopher")
	if err != nil {
		t.Error(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	res.Body.Close()

	expected := "hello, gopher"
	if expected != string(body) {
		t.Fatalf("response of /greet?name=gopher returns %s want %s", string(body), expected)
	}
}

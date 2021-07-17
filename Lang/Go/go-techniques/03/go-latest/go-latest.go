package main

import (
	"fmt"
	"github.com/tcnksm/go-latest"
)

func main() {
	html := &latest.HTMLMeta{
		URL: "http://example.com/info",
		Name: "reduce-worker",
	}

	res, _ := latest.Check(html, "0.1.0")
	if res.Outdated {
		fmt.Printf("0.1.0 is not latest, %s, upgrade to %s", res.Meta.Message, res.Current)
	}
}

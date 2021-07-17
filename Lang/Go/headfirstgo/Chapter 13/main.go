package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Url    string
	Length int
}

func GetWebpageLength(url string, ch chan Page) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	ch <- Page{url, len(body)}
}

func main() {
	ch := make(chan Page)
	urls := [3]string{
		"https://example.com",
		"https://golang.org",
		"https://golang.org/doc",
	}
	for _, url := range urls {
		go GetWebpageLength(url, ch)
	}
	for i := 0; i < len(urls); i++ {
		page := <-ch
		fmt.Printf("%s: %d\n", page.Url, page.Length)
	}
}

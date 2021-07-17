package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func View(writer http.ResponseWriter, filename string, data interface{}) error {
	html, err := template.ParseFiles(filename)
	if err != nil {
		return err
	}
	err = html.Execute(writer, data)
	if err != nil {
		return err
	}
	return nil
}

const (
	filename = "signatures.txt"
)

type GuestBook struct {
	Signatures []string
	Count      int
}

func GetStrings(filename string) ([]string, error) {
	var strings []string

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strings = append(strings, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return strings, err
	}

	return strings, nil
}

func Sign(filename string, signature string) error {
	file, err := os.OpenFile(filename, os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = fmt.Fprintf(file, "%s\n", signature)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	http.HandleFunc("/guestbook", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "GET":
			signatures, err := GetStrings(filename)
			if err != nil {
				log.Fatal(err)
			}
			guestbook := GuestBook{
				signatures, len(signatures),
			}
			if err := View(writer, "guestbook.html", guestbook); err != nil {
				log.Fatal(err)
			}
		case "POST":
			signature := request.PostFormValue("signature")
			err := Sign(filename, signature)
			if err != nil {
				log.Fatal(err)
			}
			http.Redirect(writer, request, "/guestbook", 302)
		}
	})
	http.HandleFunc("/guestbook/new", func(writer http.ResponseWriter, request *http.Request) {
		if err := View(writer, "new.html", nil); err != nil {
			log.Fatal(err)
		}
	})
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}

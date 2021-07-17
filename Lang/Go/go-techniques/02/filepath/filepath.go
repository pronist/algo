package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

func main() {
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if ok, err := path.Match("/data/*.png", request.URL.Path); !ok || err != nil {
			http.NotFound(writer, request)
			return
		}

		p := filepath.Join(cwd, "data", filepath.Base(request.URL.Path))
		file, err := os.Open(p)
		if err != nil {
			http.NotFound(writer, request)
			return
		}

		defer file.Close()
		io.Copy(writer, file)
	})

	http.HandleFunc("/upload", func(writer http.ResponseWriter, request *http.Request) {
		switch request.Method {
		case "POST":
			stream, header, err := request.FormFile("upload")
			if err != nil {
				//log.Fatal(err)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			//f, err := header.Open()

			p := filepath.Join("data", filepath.Base(header.Filename))
			f, err := os.Create(p)
			if err != nil {
				//log.Fatal(err)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			//buf, err := ioutil.ReadAll(f)
			//ioutil.WriteFile(p, buf, 0644)
			defer f.Close()
			io.Copy(f, stream)
			http.Redirect(writer, request, path.Join("/data", filepath.Base(p)), 301)
		case "GET":
			html, err := template.ParseFiles("upload.html")
			if err != nil {
				log.Fatal(err)
			}
			html.Execute(writer, nil)
		}
	})
	http.ListenAndServe(":8080", nil)
}

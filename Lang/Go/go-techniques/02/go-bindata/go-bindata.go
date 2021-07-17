package main

import (
	"bytes"
	"io"
	"mime"
	"net/http"
	"path"
	"strconv"
)

//go:generate go-bindata -prefix data ./data

func main() {
	http.Handle("/", http.StripPrefix("/", http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		p := request.URL.Path
		if p == "" {
			p = "index.html"
		}

		bs, err := Asset(p)
		if err != nil {
			http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		if ctype := mime.TypeByExtension(path.Ext(p)); ctype != "" {
			writer.Header().Set("Content-Type", ctype)
		}
		writer.Header().Set("Content-Length", strconv.Itoa(len(bs)))
		io.Copy(writer, bytes.NewBuffer(bs))
	})))

	http.ListenAndServe(":8080", nil)
}



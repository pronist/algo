package main

import (
	"os"
	"text/template"
)

func main() {
	tmpl, _ := template.New("test").Parse("Count: {{.Count}}, Name: {{.Name}}")
	tmpl.Execute(
		os.Stdout,
		struct {
			Name  string
			Count int
		}{"Hello, Go!", 10})
}

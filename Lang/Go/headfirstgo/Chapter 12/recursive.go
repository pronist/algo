package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

func reportPanic() {
	p := recover()
	if p == nil {
		return
	}
	err, ok := p.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(err)
	}
}

func ReadDir(dirname string) {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filePath := filepath.Join(dirname, file.Name())
		if file.IsDir() {
			ReadDir(filePath)
		}
		fmt.Println(filePath)
	}
}

func main() {
	defer reportPanic()
	ReadDir("./")
}

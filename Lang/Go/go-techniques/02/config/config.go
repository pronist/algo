package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

type author struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type config struct {
	Addr   string `json:"adder"`
	Author author `json:"author"`
}

func main() {
	home := os.Getenv("HOME")
	if home == "" && runtime.GOOS == "windows" {
		home = os.Getenv("USERPROFILE")
	}

	p := filepath.Join(home, ".config", "config.json")
	file, err := os.Open(p)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var cfg config
	err = json.NewDecoder(file).Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.Marshal(&cfg)
	fmt.Println(string(b))
}

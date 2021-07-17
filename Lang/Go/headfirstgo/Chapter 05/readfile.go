package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err = file.Close(); err != nil {
		log.Fatal(err)
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

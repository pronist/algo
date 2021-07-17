package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())

	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")
	// fmt.Println(target)

	for i := 0; i < 10; i++ {
		fmt.Print("Make a guess: ")

		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		guess, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			log.Fatal(err)
		}

		if guess == target {
			fmt.Println("Good job! You guessed it!")
			return
		} else if guess < target {
			fmt.Println("Oops. Your guess was LOW.")
		} else {
			fmt.Println("Oops. Your guess was HIGH.")
		}
	}
	fmt.Println("Sorry. You didn't guess number. It was:", target)
}

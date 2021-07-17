package main

import (
	"fmt"
	"time"
)

func main() {
	//var wg sync.WaitGroup

	//wg.Add(1)
	go func() {
		//defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Printf("%c\n", 'a')
		}
	}()

	//wg.Add(1)
	go func() {
		//wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Printf("%c\n", 'b')
		}
	}()

	//wg.Wait()
	time.Sleep(time.Second)
	fmt.Println("end main()")
}

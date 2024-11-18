package main

import "fmt"

func main() {
	fmt.Println("main() started")
	c := make(chan string)
	var num = 12
	// launch anonymous goroutine
	go func(c chan string) {
		fmt.Println("Hello " + <-c + "!")
	}(c)

	c <- "John"
	fmt.Println("main() stopped")
}

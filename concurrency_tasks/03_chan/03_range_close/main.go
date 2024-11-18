package main

import (
	"fmt"
	// "time"
)

// Что здесь происходит?
// Как можно сократить запись? Какой синтаксический сахар можно использовать здесь?

func squares(c chan int) {
	for i := 0; i <= 9; i++ {
		// time.Sleep(time.Millisecond)
		fmt.Println("squares", i)
		c <- i * i
	}
	close(c) // close channel
}

func main() {
	fmt.Println("main() started")
	c := make(chan int)

	go squares(c) // start goroutine

	// periodic block/unblock of main goroutine until chanel closes
	for {
		val, ok := <-c
		if ok == false {
			fmt.Println(val, ok, "<-- loop broke!")
			break // exit break loop
		} else {
			fmt.Println(val, ok)
		}
	}

	fmt.Println("main() stopped")
}

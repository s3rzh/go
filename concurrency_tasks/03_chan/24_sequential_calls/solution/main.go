package main

import "time"

func main() {
	timeStart := time.Now()

	c1, c2 := worker(), worker()

	_, _ = <-c1, <-c2

	println(int(time.Since(timeStart).Seconds())) // 3 sec.
}

func worker() chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 42
	}()

	return ch
}

// Output: 3 sec.

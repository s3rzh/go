package main

import "time"
// что выведет 3 или 6 и как исправить?
func main() {
	timeStart := time.Now()

	_, _ = <-worker(), <-worker()

	println(int(time.Since(timeStart).Seconds())) // 6 sec.
}

func worker() chan int {
	ch := make(chan int)

	go func() {
		time.Sleep(3 * time.Second)
		ch <- 42
	}()

	return ch
}

// Output: 6 sec.

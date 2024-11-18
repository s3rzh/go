package main

import (
	"fmt"
	"runtime"
)

// what happens here?
func main() {
	ch := make(chan int, 1)

	fmt.Println(runtime.NumGoroutine())
	ch <- 1
}

// программа просто завершится корректно тк канал буферезированый, ждать никого не нужно

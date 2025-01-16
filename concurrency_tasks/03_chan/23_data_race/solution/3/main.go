package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

func main() {

	var done atomic.Bool //  The zero value is false.
	runtime.GOMAXPROCS(1)

	go func() {
		done.Store(true)
	}()

	for !done.Load() {
	}
	fmt.Println("finished")
}

// те на атамиках, без блокировок

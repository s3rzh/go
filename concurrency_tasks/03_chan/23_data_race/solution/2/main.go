package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	mu := &sync.RWMutex{}

	runtime.GOMAXPROCS(1)

	done := false

	go func() {
		mu.Lock()
		done = true
		mu.Unlock()

	}()

	for {
		mu.RLock()
		if done {
			break
		}
		mu.RUnlock()
	}
	fmt.Println("finished")

}

// или если нам важно выполнять что-то в цикле - то можно через мьютекс (причём RWMutex тк у нас будет много чтений и только одна запись)

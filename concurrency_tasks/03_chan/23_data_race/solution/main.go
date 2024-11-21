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

// иле через RWMutex (много чтений, мало записей, у нас одна) это если нам важно чтобы выполнятся цикл

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var flag int32
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		time.Sleep(time.Second)
		atomic.StoreInt32(&flag, 1)
		wg.Done()
	}()

	go func() {
		for atomic.LoadInt32(&flag) == 0 {
			// крутимся в цикле, до тех пор, пока условие не станет ложным (те пока не установим флаг в 1 и выйдем из цикла).
		}
		fmt.Println("Flag set!")
		wg.Done()
	}()

	wg.Wait()
}

// safely

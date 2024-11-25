package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var rwm sync.RWMutex
	var count int

	size := 10
	queue := make([]int, size)

	wg.Add(size)
	go func() {
		for i := 0; i < size; i++ {
			go func() {
				defer wg.Done()
				rwm.Lock() // this locked region is safe for read and write operations.
				count++
				rwm.Unlock()
			}()
		}
	}()

	wg.Add(size)
	go func() {
		for i := 0; i < size; i++ {
			indx := i // for go ver. less than 1.22
			go func() {
				defer wg.Done()
				rwm.RLock() // this region is safe for concurrent reads, but not safe for concurrent writes.
				queue[indx] = count
				rwm.RUnlock()
			}()
		}
	}()

	wg.Wait()
	fmt.Println(queue)
}

// Вывод будет разный
// [8 7 8 8 6 9 9 9 9 8]
// [0 0 0 0 0 10 0 0 0 0]
// [3 0 0 0 0 0 0 0 0 0]

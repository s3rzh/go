package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	input := []int{1, 2, 3, 4, 5}

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}

	result := []int{}

	for _, num := range input {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			mu.Lock()
			time.Sleep(150 * time.Millisecond)
			data := num * 2
			result = append(result, data)
			mu.Unlock()
		}(num)
	}

	wg.Wait()

	fmt.Println(result)
}

// solution with mutex, but now we have locks in our concurrency

package main

import (
	"fmt"
	"sync"
	"time"
)

// What's the problem here?

func main() {
	input := []int{1, 2, 3, 4, 5}

	wg := &sync.WaitGroup{}
	result := []int{}

	for _, num := range input {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			time.Sleep(150 * time.Millisecond)
			data := num * 2
			result = append(result, data)

		}(num)
	}

	wg.Wait()

	fmt.Println(result)
}

// Here we'he DATA RACE
// go run -race ./01_concurrent_append/main.go

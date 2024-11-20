package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	input := []int{1, 2, 3, 4, 5}

	wg := &sync.WaitGroup{}

	result := make([]int, len(input))

	for i, num := range input {
		wg.Add(1)
		go func(num, i int) {
			defer wg.Done()
			time.Sleep(150 * time.Millisecond)
			data := num * 2
			result[i] = data
		}(num, i)
	}

	wg.Wait()

	fmt.Println(result)
}

// good solution, If we pre-allocate our slice variables (line: 14), we can write to each of them concurrently by index.

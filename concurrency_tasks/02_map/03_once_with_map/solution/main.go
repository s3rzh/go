package main

import (
	"fmt"
	"math/rand"
	"sync"
)

// selected only unique values
// find bug with reads-writes
func main() {
	alreadyStored := make(map[int]struct{})
	mu := &sync.Mutex{}
	capacity := 1000

	doubles := make([]int, 0, capacity)
	for i := 0; i < capacity; i++ {
		doubles = append(doubles, rand.Intn(10)) // create rand num 0...9
	}
	// 3, 4, 5, 0, 4, 9, 8, 6, 6, 5, 5, 4, 4, 4, 2, 1, 2, 3, 1 ...

	uniqueIDs := make(chan int, capacity) // канал для уникальных значений
	wg := &sync.WaitGroup{}

	for i := 0; i < capacity; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock() // ставим блокировку и для чтения (32) и для записи (33)
			defer mu.Unlock()
			if _, ok := alreadyStored[doubles[i]]; !ok {
				alreadyStored[doubles[i]] = struct{}{}
	
				uniqueIDs <- doubles[i]
			}
		}()
	}

	wg.Wait()
	close(uniqueIDs) // не забываем закрыть канал, чтобы for range ниже завершился
	for val := range uniqueIDs {
		fmt.Println(val)
	}
	// close(uniqueIDs) // здесь неправильно
	fmt.Printf("len of ids: %d\n", len(uniqueIDs)) // 0
	fmt.Println(uniqueIDs)
}

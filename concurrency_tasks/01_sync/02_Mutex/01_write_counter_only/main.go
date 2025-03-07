package main

import (
	"fmt"
	"sync"
)

var counter int
var mu sync.Mutex

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			fmt.Println("Значение счетчика:", counter)
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("Все горутины завершили работу")
}

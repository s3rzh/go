package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	input := []int{1, 2, 3, 4, 5}

	wg := &sync.WaitGroup{}
	resultChan := make(chan int, len(input))
	result := make([]int, 0, len(input))

	for _, num := range input {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			time.Sleep(150 * time.Millisecond)
			data := num * 2
			resultChan <- data
		}(num)
	}

  // вариант 1 c закрытием канала и for-range:
	//go func() {
	//	wg.Wait()
	//	close(resultChan)
	//}()

  // вариант 2 с проверкой кол-ва элементов в канале с вычитываение до 0:
	wg.Wait()

   // вариант 1:
  //for data := range resultChan {
	//	result = append(result, data)
	//}

  // вариант 2:
	for len(resultChan) > 0 {
		data := <-resultChan
		result = append(result, data) // результат будет неупорядочен, например [8 4 10 6 2]
	}

	fmt.Println(result)
}

// можно через канал, те они потокобезопасные (thread-safe), но мы теряем порядок, а возможно это важно.

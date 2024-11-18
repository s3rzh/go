package main

import (
	"fmt"
	"sync"
)

// merge two channels
func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)

	ch1 <- 1
	ch2 <- 2
	ch2 <- 4
	close(ch1)
	close(ch2)

	ch3 := syncMerge[int](ch1, ch2)

	for val := range ch3 { 
		fmt.Println(val)
	}
}

func syncMerge[T any](chans ...chan T) chan T {
	res := make(chan T)
	wg := sync.WaitGroup{}
	wg.Add(len(chans))
	for _, ch := range chans {
		go func(ch chan T) {
			defer wg.Done()
			for c := range ch {
				res <- c
			}
		}(ch)
	}
		
  wg.Wait()
	close(res)

	return res
}

package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

\
func main() {
	ch := make(chan int)
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
  	defer cancel() // на всякий случай, чтобы утечки зависшей горутине не было
  
	go func() {
		select {
		case val := <-ch:
			fmt.Println(val)
		case <-ctx.Done():
			fmt.Println("Done")
		}
	}()
	ch <- 42 // или cancel()
	time.Sleep(time.Second)
	fmt.Println(runtime.NumGoroutine()) // показываем кол-во активных горутин в момент этого вызова
}

// здесь мы добавили контекст
// и записали значение в канал

// по сути можно делать что-то одно
// отменять контекст
// или
// записывать значение в канал

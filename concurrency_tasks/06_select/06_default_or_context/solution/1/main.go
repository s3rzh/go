package main

import (
	"fmt"
	"runtime"
	"time"
)


func main() {
	ch := make(chan int)

	go func() {
		select {
		case val := <-ch:
			fmt.Println(val)
		default:                  
			fmt.Println("default")
		}
	}()
	time.Sleep(time.Second)
	fmt.Println(runtime.NumGoroutine()) // показываем кол-во активных горутин в момент этого вызова
}

// решение - добавление default сценария

// добавление горутины, чтобы увидеть количество активных горутин
// при этом, если мы уберем default на 17 и 18 строке, у нас произойдет утечка горутины
// main завершит свою работу, а созданная нами горутина нет

// также можно решить задачу, используя контекст (см 2)

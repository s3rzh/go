package main

import (
	"fmt"
	"runtime"
)

// Какая тут проблема и как решить?

func main() {
   runtime.GOMAXPROCS(1)

   done := false

   go func() {
      done = true
   }()

   for !done {
   }
   fmt.Println("finished")
}

// У нас тут есть DATA RACE несмотря на то, что главная горутина и саб-горутина (строка 14) будут выполнятся на одном P (тк это задано строкой runtime.GOMAXPROCS(1))
// в какой-то момент планировщик таки переключит контект на саб-горутину, тк поймёт, что в цикле ничего не происходит.

// Data race это состояние когда разные потоки обращаются к одной ячейке памяти без какой-либо синхронизации и как минимум один из потоков осуществляет запись (те это происходит на строках 15 запись и 18 чтение).
// GOMAXPROCS устанавливает количество задействованых P в runtime (см модель G-M-P), которые участвуют в выполнение G и никак не влияют на кол-во доступных M (Thread). M могут добавляться и удаляться при необходимости

// A typical fix for DATA RACE is to use a channel or a mutex. To preserve the lock-free behavior, one can also use the sync/atomic package.

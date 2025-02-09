
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

// У нас тут есть DATA RACE несмотря на то, что главная горутина и саб-горутина (строка 14) будут выполнятся на одном P (тк это задано строкой runtime.GOMAXPROCS(1)). Возможно это из-а НЕАТАМАРНОСТИ ОПЕРАЦИЙ (на 16ой строке, например)?
// в какой-то момент планировщик таки переключит контект на саб-горутину, тк поймёт, что в цикле ничего не происходит.

// Data race это состояние когда разные потоки обращаются к одной ячейке памяти без какой-либо синхронизации и как минимум один из потоков осуществляет запись (те это происходит на строках 15 запись и 18 чтение).
// GOMAXPROCS устанавливает количество задействованых P в runtime (см модель G-M-P), которые участвуют в выполнение G и никак не влияют на кол-во доступных M (Thread). M могут добавляться и удаляться при необходимости

// A typical fix for DATA RACE is to use a channel or a mutex. To preserve the lock-free behavior, one can also use the sync/atomic package.

// Компелятор GO, да и процессор может переставлять инструкции (для оптимизации). Те они не обязательно будут выполнятся в той последовательности, котрой написал программист. (см тут https://vkvideo.ru/video-139172865_456239228)
// Для того, чтобы запретить такое перемещание - есть Барьеры памяти (они влияюст только на порядок и только).
// те между типа операций (а это чтение и запись и их комбинации, в том числе и с самими собой, например чтение-чтение) поставить этот барьер, который гарантирует, что все чтения гарантировано будут выполнены до операций записи.
// Например Полный барьер запрешает компилятору (и процесоору) передвигать любые операции.

// Когда мы используем пакет atomic - мы под капотом использует full brarieer (полный барье) и инструкции не перепрыгивают.

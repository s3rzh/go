package main

import (
	"fmt"
	"runtime"
)

func main() {

	ch := make(chan struct{})

	runtime.GOMAXPROCS(1)

	done := false

	go func() {
		done = true
		ch <- struct{}{}
	}()

	<-ch
	for !done {

	}

	fmt.Println("finished")
}

// решение, нам придётся упорядочить наши горутины, при помощи небуферизированого канала.
// те на 21 строке мы блочим основную горутину и перелючаем на саб-горутину, которая после выполняние посылает сигнал в канал и затем заблокируеся основная горутина.

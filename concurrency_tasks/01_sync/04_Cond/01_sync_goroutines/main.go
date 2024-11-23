package main

import (
	"fmt"
	"sync"
	"time"
)

// sync.Cond можно использовать для синхронизации выполнения нескольких горутин. Например, у вас могут быть различные горутины, которые должны дождаться выполнения определенного условия, прежде чем продолжить.

func main() {
	var wg sync.WaitGroup

	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	wg.Add(2)

	go func() {
		fmt.Println("Goroutine 1 is started")
		defer wg.Done()
		cond.L.Lock() // получает блокировку (но позже её вынужен снять, а после сигнала опять захватит её)
		defer cond.L.Unlock()
		fmt.Println("Goroutine 1 is waiting for condition")
		cond.Wait() // но тк условие еще не выполнено (сигнал еще никто не посылал), данная горутина снимает блокировку и приостанавливает свое выполнение.
		fmt.Println("Goroutine 1 met the condition")
		fmt.Println("Goroutine 1 is done")
	}()

	go func() {
		fmt.Println("Goroutine 2 is started")
		defer wg.Done()
		time.Sleep(3 * time.Second) // имитация какой-то работы
		cond.L.Lock() // получает блокировку
		defer cond.L.Unlock()
		fmt.Println("Goroutine 2 is signaling condition")
		cond.Signal() // отправляет сигнал и пробуждает ожидающую горутину, которая затем получает блокировку и выполняется, но если сигнала ждет более одной горутины - то будет выбрана одна произвольно
		fmt.Println("Goroutine 2 completed signaling")
		fmt.Println("Goroutine 2 is done")
	}()

	wg.Wait()
}

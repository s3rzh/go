package main

import (
	"fmt"
	"sync"
)

// Проинициализировать строго один раз

type Singleton struct {
	data string
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() { // есть единственный метод Do
		fmt.Println("Creating Singleton instance")
		instance = &Singleton{data: "I'm the only one!"}
	})
	return instance
}

func main() {
	const N = 5

	wg := &sync.WaitGroup{}

	wg.Add(N)

	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			fmt.Printf("%p\n", GetInstance())
		}()
	}

	// Wait for goroutines to finish
	wg.Wait()
}

// В результате, адрес instance будет один и тот же (выведен 5 раз на строке 33)

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Map

	wg.Add(5)
	go func() {
		for i := 0; i < 5; i++ {
			//wg.Add(1)
			go func(j int) {
				defer wg.Done()
				m.Store(j, fmt.Sprintf("%v", j))
			}(i)
		}
	}()

	wg.Add(1000)
	go func() {
		for i := 0; i < 1000; i++ {
			//wg.Add(1)
			go func() {
				defer wg.Done()
				key := rand.Intn(5) // генерим наши ключи от 0 до 5 (невключительно)
				v, ok := m.Load(key)
				fmt.Println("for loop: ", key, v, ok)
			}()
		}
	}()

	wg.Wait()
	fmt.Println("Done.")
}

// Если раскоментить 16 и 27 строку и выполнить с флагом -race то переодически будет возникать DATA RACE на 16 (wg.Add(1)) и 37 (wg.Wait()) строке, почему?

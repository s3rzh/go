package main

import (
	"fmt"
)

// Сделать кастомную waitGroup на семафоре.

type sema chan struct{}

func New(n int) sema {
	return make(sema, n)
}

func (s sema) Inc(k int) {
	for i := 0; i < k; i++ {
		s <- struct{}{}
	}
}

func (s sema) Dec(k int) {
	for i := 0; i < k; i++ {
		<-s
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	n := len(numbers)

	sem := New(n)

	for _, num := range numbers {
		go func(n int) {
			fmt.Println(n)
			sem.Inc(1)
		}(num)
	}

	sem.Dec(n)

}

// Семафор можно легко получить из канала. Чтоб не аллоцировать лишние данные, будем складывать туда пустые структуры.
// В нашем случае мы хотим сделать семафор, который будет ждать выполнения пяти горутин.

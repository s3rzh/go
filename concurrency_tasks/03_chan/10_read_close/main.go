package main

import "fmt"

// что выведет программа?
// можно ли читать данные из закрытого канала?

func main() {
	c := make(chan int, 3)
	c <- 1
	c <- 2
	c <- 3
	close(c)

	// iteration terminates after receiving 3 values
	for elem := range c {
		fmt.Println(elem)
	}
	val, ok := <-c
	fmt.Println(val, ok) // 0 false (те читать из закрытого канала мы можем и выведет нулевое значение, для int это 0)
}

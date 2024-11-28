package main

import "fmt"

func fibonacci(n uint) uint {
	if n < 2 {
		return n
	}

	var a, b uint

	b = 1

	for n--; n > 0; n-- {
		a += b
		a, b = b, a
	}

	return b
}

func main() {
	fmt.Println(fibonacci(1_000))
}

// Решение без рекурсии

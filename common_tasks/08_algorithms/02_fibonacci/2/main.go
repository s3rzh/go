package main

import "fmt"

var (
	fibonacciCache = make(map[uint]uint)
)

func fibonacci(n uint) uint {
	if n < 2 {
		return n
	}

	if result, ok := fibonacciCache[n]; ok {
		return result
	}

	result := fibonacci(n-1) + fibonacci(n-2)

	fibonacciCache[n] = result

	return result
}

func main() {
	fmt.Println(fibonacci(1_000)) // те тысячное число фибоначчи равно 817770325994397771
}

// Для оптимизации задействуем кэш, в котором будем хранить, ранее вычисленые значения.

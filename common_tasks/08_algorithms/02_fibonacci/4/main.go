package main

import (
	"fmt"
	"math/big"
)

func fibonacci(n uint) *big.Int {
	if n < 2 {
		return big.NewInt(int64(n))
	}

	a, b := big.NewInt(0), big.NewInt(1)

	for n--; n > 0; n-- {
		a.Add(a, b)
		a, b = b, a
	}

	return b
}

func main() {
	fmt.Println(fibonacci(5_000))
}

// Вариант функции для очень больших числе фибоначчи.

package main

import (
	"fmt"
)

func main() {
	test1 := []int{1, 2, 3, 4, 5}
	test1 = test1[:3]  // [1 2 3] 3 5
	test2 := test1[3:] //test2: [] пусто и len=0 cap=2
	fmt.Println(test2)

	// но тоже самое, реслайсинг с указаниме 5
	test3 := test1[3:5] //test3: [4 5] и len=2 cap=2
	fmt.Println(test3)
}

// Найти почему есть разница между test2 и test3?
